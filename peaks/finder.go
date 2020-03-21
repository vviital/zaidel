package peaks

import (
	"errors"
	"log"
	"math"

	"github.com/vviital/zaidel/geometry"
)

const fMaxPeaks = 100
const bw = 2
const minThreshold = 0
const maxThreshold = 100
const peakWindow = 1024
const epsilon = 0.000000001

func calculateBackground(source []float64, workingSpace []float64, l1low float64, shift int, settings searchSettings) (background []float64) {
	for i := 1; i <= settings.NumberIterations; i++ {
		for j := i; j < settings.SizeExtended-i; j++ {
			if !settings.SmoothMarkov {
				a := workingSpace[settings.SizeExtended+j]
				b := (workingSpace[settings.SizeExtended+j-i] + workingSpace[settings.SizeExtended+j+i]) / 2.0
				if b < a {
					a = b
				}
				workingSpace[j] = a
			} else {
				a := workingSpace[settings.SizeExtended+j]
				var av, men, b, c float64
				var w int

				for w = j - bw; w <= j+bw; w++ {
					if w >= 0 && w < settings.SizeExtended {
						av += workingSpace[settings.SizeExtended+w]
						men++
					}
				}
				av = av / men
				men = 0
				for w = j - i - bw; w <= j-i+bw; w++ {
					if w >= 0 && w < settings.SizeExtended {
						b += workingSpace[settings.SizeExtended+w]
						men++
					}
				}
				b = b / men
				men = 0
				for w = j + i - bw; w <= j+i+bw; w++ {
					if w >= 0 && w < settings.SizeExtended {
						c += workingSpace[settings.SizeExtended+w]
						men++
					}
				}
				c = c / men
				b = (b + c) / 2
				if b < a {
					av = b
				}
				workingSpace[j] = av
			}
		}
		for j := i; j < settings.SizeExtended-i; j++ {
			workingSpace[settings.SizeExtended+j] = workingSpace[j]
		}
	}
	for j := 0; j < settings.SizeExtended; j++ {
		if j < shift {
			a := float64(j - shift)
			b := source[0] + l1low*a
			if b < 0 {
				b = 0
			}
			workingSpace[settings.SizeExtended+j] = b - workingSpace[settings.SizeExtended+j]
		} else if j >= settings.Size+shift {
			// a := float64(j - (settings.Size - 1 + shift))
			b := source[settings.Size-1]
			if b < 0 {
				b = 0
			}
			workingSpace[settings.SizeExtended+j] = b - workingSpace[settings.SizeExtended+j]
		} else {
			workingSpace[settings.SizeExtended+j] = source[j-shift] - workingSpace[settings.SizeExtended+j]
		}
	}
	for j := 0; j < settings.SizeExtended; j++ {
		if workingSpace[settings.SizeExtended+j] < 0 {
			workingSpace[settings.SizeExtended+j] = 0
		}
	}

	background = make([]float64, settings.Size)
	copy(background, workingSpace[(settings.SizeExtended-settings.Size)/2:(settings.SizeExtended-settings.Size)/2+settings.Size])
	return
}

func getL1Low(source []float64, k int) float64 {
	var m0low, m1low, m2low float64
	var l0low, l1low float64

	if k >= 2 {
		for i := 0; i < k; i++ {
			a := float64(i)
			b := source[i]
			m0low += 1
			m1low += a
			m2low += a * a
			l0low += b
			l1low += a * b
		}
		detlow := m0low*m2low - m1low*m1low
		if detlow != 0 {
			l1low = (-l0low*m1low + l1low*m0low) / detlow
		} else {
			l1low = 0
		}

		if l1low > 0 {
			l1low = 0
		}
	} else {
		l1low = 0
	}

	return l1low
}

func createWorkingSpace(source []float64, l1low float64, settings searchSettings) (workingSpace []float64) {
	workingSpace = make([]float64, 7*(settings.Size+2*(int)(7*settings.Sigma+0.5)))
	var shift = settings.NumberIterations

	for i := 0; i < settings.SizeExtended; i++ {
		if i < shift {
			a := float64(i - shift)
			workingSpace[i+settings.SizeExtended] = source[0] + l1low*a
			if workingSpace[i+settings.SizeExtended] < 0 {
				workingSpace[i+settings.SizeExtended] = 0
			}
		} else if i >= settings.Size+shift {
			workingSpace[i+settings.SizeExtended] = source[settings.Size-1]
			if workingSpace[i+settings.SizeExtended] < 0 {
				workingSpace[i+settings.SizeExtended] = 0
			}
		} else {
			workingSpace[i+settings.SizeExtended] = source[i-shift]
		}
	}

	return workingSpace
}

func generateResponseVector(workingSpace []float64, settings searchSettings) (lh_gold int, area float64, posit int) {
	var maximum float64
	lh_gold = -1

	//generate response vector
	for i := 0; i < settings.SizeExtended; i++ {
		lda := float64(i) - 3*settings.Sigma
		lda = lda * lda / (2 * settings.Sigma * settings.Sigma)
		j := (int)(1000 * math.Exp(-lda))
		lda = float64(j)
		if lda != 0 {
			lh_gold = i + 1
		}
		workingSpace[i] = lda
		area = area + lda
		if lda > maximum {
			maximum = lda
			posit = i
		}
	}

	return
}

func createMatrix(workingSpace []float64, lh_gold int, settings searchSettings) {
	//create matrix at*a(vector b)
	i := lh_gold - 1
	if i > settings.SizeExtended {
		i = settings.SizeExtended
	}
	imin := -i
	imax := i
	for i := imin; i <= imax; i++ {
		var lda, ldb, ldc float64

		jmin := 0
		if i < 0 {
			jmin = -i
		}

		jmax := lh_gold - 1 - i
		if jmax > (lh_gold - 1) {
			jmax = lh_gold - 1
		}

		for j := jmin; j <= jmax; j++ {
			ldb = workingSpace[j]
			ldc = workingSpace[i+j]
			lda = lda + ldb*ldc
		}
		workingSpace[settings.SizeExtended+i-imin] = lda
	}
}

func createVectorP(workingSpace []float64, lh_gold int, settings searchSettings) {
	//create vector p
	i := lh_gold - 1
	imin := -i
	imax := settings.SizeExtended + i - 1

	for i = imin; i <= imax; i++ {
		lda := float64(0)

		for j := 0; j <= (lh_gold - 1); j++ {
			ldb := workingSpace[j]
			k := i + j
			if k >= 0 && k < settings.SizeExtended {
				ldc := workingSpace[2*settings.SizeExtended+k]
				lda = lda + ldb*ldc
			}
		}
		workingSpace[4*settings.SizeExtended+i-imin] = lda
	}
	//move vector p
	for i = imin; i <= imax; i++ {
		workingSpace[2*settings.SizeExtended+i-imin] = workingSpace[4*settings.SizeExtended+i-imin]
	}
}

func deconvolution(workingSpace []float64, settings searchSettings) (float64, float64) {
	//deconvolution starts
	shift := settings.NumberIterations
	lh_gold, area, posit := generateResponseVector(workingSpace, settings)

	//read source vector
	for i := 0; i < settings.SizeExtended; i++ {
		workingSpace[2*settings.SizeExtended+i] = math.Abs(workingSpace[settings.SizeExtended+i])
	}

	createMatrix(workingSpace, lh_gold, settings)
	createVectorP(workingSpace, lh_gold, settings)

	//initialization of resulting vector
	for i := 0; i < settings.SizeExtended; i++ {
		workingSpace[i] = 1
	}

	//START OF ITERATIONS
	for lindex := 0; lindex < settings.DeconvolutionIterations; lindex++ {
		for i := 0; i < settings.SizeExtended; i++ {
			if math.Abs(workingSpace[2*settings.SizeExtended+i]) > epsilon && math.Abs(workingSpace[i]) > epsilon {
				lda := float64(0)

				jmin := lh_gold - 1
				if jmin > i {
					jmin = i
				}

				jmin = -jmin
				jmax := lh_gold - 1
				if jmax > (settings.SizeExtended - 1 - i) {
					jmax = settings.SizeExtended - 1 - i
				}

				for j := jmin; j <= jmax; j++ {
					ldb := workingSpace[j+lh_gold-1+settings.SizeExtended]
					ldc := workingSpace[i+j]
					lda = lda + ldb*ldc
				}
				ldb := workingSpace[2*settings.SizeExtended+i]
				if lda != 0 {
					lda = ldb / lda
				} else {
					lda = 0
				}

				ldb = workingSpace[i]
				lda = lda * ldb
				workingSpace[3*settings.SizeExtended+i] = lda
			}
		}
		for i := 0; i < settings.SizeExtended; i++ {
			workingSpace[i] = workingSpace[3*settings.SizeExtended+i]
		}
	}
	//shift resulting spectrum
	for i := 0; i < settings.SizeExtended; i++ {
		lda := workingSpace[i]
		j := i + posit
		j = j % settings.SizeExtended
		workingSpace[settings.SizeExtended+j] = lda
	}
	//write back resulting spectrum
	var maximum, maximum_decon float64
	j := lh_gold - 1
	for i := 0; i < settings.SizeExtended-j; i++ {
		if i >= shift && i < settings.Size+shift {
			workingSpace[i] = area * workingSpace[settings.SizeExtended+i+j]
			if maximum_decon < workingSpace[i] {
				maximum_decon = workingSpace[i]
			}

			if maximum < workingSpace[6*settings.SizeExtended+i] {
				maximum = workingSpace[6*settings.SizeExtended+i]
			}
		} else {
			workingSpace[i] = 0
		}
	}

	return maximum, maximum_decon
}

func findPeaksInTheSpectrum(workingSpace []float64, lda, maximum_decon, maximum float64, settings searchSettings) (peakIndex int, fPositionX []float64) {
	shift := settings.NumberIterations
	fPositionX = make([]float64, 100)

	//searching for peaks in deconvolved spectrum
	for i := 1; i < settings.SizeExtended-1; i++ {
		if workingSpace[i] > workingSpace[i-1] && workingSpace[i] > workingSpace[i+1] {
			if i >= shift && i < settings.Size+shift {
				if workingSpace[i] > lda*maximum_decon && workingSpace[6*settings.SizeExtended+i] > settings.Threshold*maximum/100.0 {
					var a, b float64

					for j := i - 1; j <= i+1; j++ {
						a += float64((j - shift)) * workingSpace[j]
						b += workingSpace[j]
					}
					a = a / b
					if a < 0 {
						a = 0
					}
					if a >= float64(settings.Size) {
						a = float64(settings.Size - 1)
					}

					if peakIndex == 0 {
						fPositionX[0] = a
						peakIndex = 1
					} else {
						priz := 0
						j := 0

						for ; j < peakIndex && priz == 0; j++ {
							if workingSpace[6*settings.SizeExtended+shift+int(a)] >
								workingSpace[6*settings.SizeExtended+shift+int(fPositionX[j])] {
								priz = 1
							}
						}
						if priz == 0 && j < fMaxPeaks {
							fPositionX[j] = a
						} else if priz != 0 {
							for k := peakIndex; k >= j; k-- {
								if k < fMaxPeaks {
									fPositionX[k] = fPositionX[k-1]
								}
							}
							fPositionX[j-1] = a
						}
						if peakIndex < fMaxPeaks {
							peakIndex += 1
						}
					}
				}
			}
		}
	}

	if peakIndex == fMaxPeaks {
		log.Print("Peaks out of buffer.")
	}

	return
}

func getLDA(settings searchSettings) (lda float64) {
	lda = 1
	if lda > settings.Threshold {
		lda = settings.Threshold
	}
	lda = lda / 100
	return
}

// SearchHighRes function provides users with the possibility to vary the
// input parameters and with the access to the output deconvolved data in the
// destination spectrum. Based on the output data one can tune the parameters.
// Please see original implementation: https://root.cern.ch/doc/v608/TSpectrum_8cxx_source.html
func SearchHighRes(source []float64, settings searchSettings) (fPositionX, background []float64, peakIndex int, err error) {
	var a, b float64
	var shift = settings.NumberIterations
	l1low := getL1Low(source, (int)(2*settings.Sigma+0.5))
	workingSpace := createWorkingSpace(source, l1low, settings)

	if settings.CalculateBackground {
		background = calculateBackground(source, workingSpace, l1low, shift, settings)
	}
	for i := 0; i < settings.SizeExtended; i++ {
		workingSpace[i+6*settings.SizeExtended] = workingSpace[i+settings.SizeExtended]
	}
	if settings.SmoothMarkov {
		for j := 0; j < settings.SizeExtended; j++ {
			workingSpace[2*settings.SizeExtended+j] = workingSpace[settings.SizeExtended+j]
		}

		settingsCopy := settings
		settingsCopy.Size = settings.SizeExtended
		SmoothMarkov(workingSpace[2*settings.SizeExtended:2*settings.SizeExtended+settings.SizeExtended], settingsCopy)

		for j := 0; j < settings.SizeExtended; j++ {
			workingSpace[settings.SizeExtended+j] = workingSpace[2*settings.SizeExtended+j]
		}
		if settings.CalculateBackground == true {
			for i := 1; i <= settings.NumberIterations; i++ {
				for j := i; j < settings.SizeExtended-i; j++ {
					a = workingSpace[settings.SizeExtended+j]
					b = (workingSpace[settings.SizeExtended+j-i] + workingSpace[settings.SizeExtended+j+i]) / 2.0
					if b < a {
						a = b
					}

					workingSpace[j] = a
				}
				for j := i; j < settings.SizeExtended-i; j++ {
					workingSpace[settings.SizeExtended+j] = workingSpace[j]
				}

			}
			for j := 0; j < settings.SizeExtended; j++ {
				workingSpace[settings.SizeExtended+j] = workingSpace[2*settings.SizeExtended+j] - workingSpace[settings.SizeExtended+j]
			}
		}
	}
	maximum, maximum_decon := deconvolution(workingSpace, settings)
	peakIndex, fPositionX = findPeaksInTheSpectrum(workingSpace, getLDA(settings), maximum_decon, maximum, settings)
	return
}

// SmoothMarkov calculates smoothed spectrum from source spectrum based on Markov chain method.
func SmoothMarkov(source []float64, settings searchSettings) error {
	var xmin, l int
	var a, b, maxch, nom, nip, nim, sp, sm, area float64

	workingSpace := make([]float64, settings.Size)
	xmax := settings.Size - 1

	for i := 0; i < settings.Size; i++ {
		workingSpace[i] = 0
		if maxch < source[i] {
			maxch = source[i]
		}

		area += source[i]
	}
	if maxch == 0 {
		return errors.New("invalid source value")
	}

	nom = 1
	workingSpace[xmin] = 1
	for i := xmin; i < xmax; i++ {
		nip = source[i] / maxch
		nim = source[i+1] / maxch
		sp = 0
		sm = 0
		for l = 1; l <= settings.AverageWindow; l++ {
			if (i + l) > xmax {
				a = source[xmax] / maxch
			} else {
				a = source[i+l] / maxch
			}

			b = a - nip
			if a+nip <= 0 {
				a = 1
			} else {
				a = math.Sqrt(a + nip)
			}

			b = b / a
			b = math.Exp(b)
			sp = sp + b
			if (i - l + 1) < xmin {
				a = source[xmin] / maxch
			} else {
				a = source[i-l+1] / maxch
			}

			b = a - nim
			if a+nim <= 0 {
				a = 1
			} else {
				a = math.Sqrt(a + nim)
			}

			b = b / a
			b = math.Exp(b)
			sm = sm + b
		}
		a = sp / sm
		workingSpace[i+1] = workingSpace[i] * a
		a = workingSpace[i+1]
		nom = nom + a
	}
	for i := xmin; i <= xmax; i++ {
		workingSpace[i] = workingSpace[i] / nom
	}
	for i := 0; i < settings.Size; i++ {
		source[i] = workingSpace[i] * area
	}

	return nil
}

func CombineCoordinatesWithFloats(points geometry.Coordinates, floats []float64) (pts geometry.Coordinates) {
	for i, point := range points {
		pts = append(pts, geometry.Coordinate{X: point.X, Y: floats[i]})
	}
	return
}

func IsCrossedAndIndex(firstSequence, secondSequence geometry.Coordinates, index int) bool {
	return (len(secondSequence)-1 <= index || len(firstSequence)-1 <= index) ||
		(firstSequence[index].Y == secondSequence[index].Y || (firstSequence[index-1].Y >= secondSequence[index-1].Y &&
			firstSequence[index+1].Y <= secondSequence[index+1].Y) || (firstSequence[index+1].Y >= secondSequence[index+1].Y &&
			firstSequence[index-1].Y <= secondSequence[index-1].Y))
}

// validateParams checks if params have valid values
func validateParams(settings searchSettings) error {
	if settings.Sigma < 1 || (int)(5.0*settings.Sigma+0.5) > peakWindow/2 {
		return errors.New("invalid sigma value")
	}

	if settings.Threshold <= minThreshold || settings.Threshold >= maxThreshold {
		return errors.New("invalid threshold value")
	}

	if settings.AverageWindow <= 0 {
		return errors.New("invalid averWindow value")
	}

	if settings.CalculateBackground && settings.Size < 2*settings.NumberIterations+1 {
		return errors.New("invalid sigma value")
	}

	return nil
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Find return information about peaks
func Find(points geometry.Coordinates, peakSearchSettings PeakSearchSettings) (FinderResult, error) {
	settings := newSearchSettings(points, peakSearchSettings)

	if err := validateParams(settings); err != nil {
		return FinderResult{}, err
	}

	ys := points.GetYs()

	fPositionX, background, peaksCount, err := SearchHighRes(ys, settings)

	if err != nil {
		return FinderResult{}, err
	}

	var peaks []Peak
	waveletList := ys
	err = SmoothMarkov(waveletList, settings)

	if err != nil {
		return FinderResult{}, err
	}

	waveletData := CombineCoordinatesWithFloats(points, waveletList)
	backgroundData := CombineCoordinatesWithFloats(points, background)

	for i := 0; i < peaksCount; i++ {
		// TODO: check why sometimes it's goes beyond the array size
		var binR = min(1+(int)(fPositionX[i]+0.5), len(points)-1)
		var binM = min((int)(fPositionX[i]+0.5), len(points)-1)
		var binL = min((int)(fPositionX[i]), len(points)-1)

		point := geometry.Coordinates{points[binR], points[binM], points[binL]}.GetMaxPointByY()

		var peak = Peak{Point: point}
		searchStartIndex, _ := backgroundData.FindByX(point.X)

		//left search <----
		var leftIndex = 0
		var rightIndex = 0
		for bi := searchStartIndex; bi > 0; bi-- {
			if IsCrossedAndIndex(points, backgroundData, bi) {
				leftIndex = bi
				peak.Left = points[bi]
				break
			}
		}
		//right search ---->
		for bi := searchStartIndex; bi < len(backgroundData); bi++ {
			if IsCrossedAndIndex(points, backgroundData, bi) {
				rightIndex = bi
				peak.Right = points[bi]
				break
			}
		}

		peak.Area = geometry.CalculateArea(leftIndex, rightIndex, points, backgroundData)
		peaks = append(peaks, peak)
	}

	return FinderResult{
		BackgroundData:         backgroundData,
		OriginalPeaksPositions: fPositionX,
		Peaks:                  peaks,
		PeaksCount:             peaksCount,
		SpectrumArea:           geometry.CalculateSpectrumArea(points, backgroundData),
		WaveletData:            waveletData,
	}, nil
}
