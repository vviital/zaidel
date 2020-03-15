package spectrumlines

import (
	"encoding/json"

	"github.com/vviital/zaidel/peaks"
)

var peaksOne []peaks.Peak
var spectrumLinesOne []DeterminedPeak

func init() {
	json.Unmarshal([]byte(peaksOneJSON), &peaksOne)
	json.Unmarshal([]byte(spectrumLinesOneJSON), &spectrumLinesOne)
}

var peaksOneJSON = `
[
  {
    "peak": {
      "x": 324.754,
      "y": 13518
    },
    "left": {
      "x": 324.037035019455,
      "y": 592.666666666667
    },
    "right": {
      "x": 325.354454545455,
      "y": 604.333333333333
    },
    "area": 3435.747801653668
  },
  {
    "peak": {
      "x": 327.396,
      "y": 10168
    },
    "left": {
      "x": 326.795545454545,
      "y": 632.666666666667
    },
    "right": {
      "x": 328.488535714286,
      "y": 723
    },
    "area": 2768.1999334973566
  },
  {
    "peak": {
      "x": 334.563882352941,
      "y": 8914.33333333333
    },
    "left": {
      "x": 333.546882352941,
      "y": 556
    },
    "right": {
      "x": 335.162117647059,
      "y": 574.666666666667
    },
    "area": 2208.789759337079
  },
  {
    "peak": {
      "x": 330.309428571429,
      "y": 6416.33333333333
    },
    "left": {
      "x": 329.884553571429,
      "y": 670.333333333333
    },
    "right": {
      "x": 331.034294117647,
      "y": 664.333333333333
    },
    "area": 1588.5561326854172
  },
  {
    "peak": {
      "x": 328.24575,
      "y": 3173
    },
    "left": {
      "x": 326.795545454545,
      "y": 632.666666666667
    },
    "right": {
      "x": 328.488535714286,
      "y": 723
    },
    "area": 2768.1999334973566
  },
  {
    "peak": {
      "x": 352.533024,
      "y": 2411.66666666667
    },
    "left": {
      "x": 351.937784,
      "y": 549.333333333333
    },
    "right": {
      "x": 353.544932,
      "y": 481.666666666667
    },
    "area": 452.47408993394004
  },
  {
    "peak": {
      "x": 341.503411764706,
      "y": 2069
    },
    "left": {
      "x": 341.204294117647,
      "y": 462
    },
    "right": {
      "x": 341.802529411765,
      "y": 480.333333333333
    },
    "area": 354.06871271224645
  },
  {
    "peak": {
      "x": 361.986291666667,
      "y": 2042.33333333333
    },
    "left": {
      "x": 361.808666666667,
      "y": 483
    },
    "right": {
      "x": 362.282333333333,
      "y": 497.666666666667
    },
    "area": 312.11189111913336
  },
  {
    "peak": {
      "x": 351.58064,
      "y": 1855
    },
    "left": {
      "x": 350.9854,
      "y": 656.666666666667
    },
    "right": {
      "x": 351.818736,
      "y": 518.666666666667
    },
    "area": 398.98992332736634
  }
]
`

var spectrumLinesOneJSON = `
[
  {
    "peak": {
      "x": 324.754,
      "y": 13518
    },
    "left": {
      "x": 324.037035019455,
      "y": 592.666666666667
    },
    "right": {
      "x": 325.354454545455,
      "y": 604.333333333333
    },
    "area": 3435.747801653668,
    "Elements": [
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Cu",
        "waveLength": 324.754,
        "isSearchCriteriaMatched": true,
        "similarity": 1
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Mn",
        "waveLength": 324.754,
        "isSearchCriteriaMatched": true,
        "similarity": 1
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Sb",
        "waveLength": 324.754,
        "isSearchCriteriaMatched": true,
        "similarity": 1
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ag",
        "waveLength": 324.755,
        "isSearchCriteriaMatched": true,
        "similarity": 0.9800000000004729
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 324.755,
        "isSearchCriteriaMatched": true,
        "similarity": 0.9800000000004729
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Nb",
        "waveLength": 324.748,
        "isSearchCriteriaMatched": true,
        "similarity": 0.879999999999427
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "In",
        "waveLength": 324.761,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8599999999999
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Tm",
        "waveLength": 324.747,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8599999999999
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Mo",
        "waveLength": 324.762,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8400000000003729
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Hf",
        "waveLength": 324.767,
        "isSearchCriteriaMatched": true,
        "similarity": 0.7400000000004638
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Fe",
        "waveLength": 324.739,
        "isSearchCriteriaMatched": true,
        "similarity": 0.699999999999136
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Fe",
        "waveLength": 324.728,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4799999999997908
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 324.726,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4399999999995998
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Cr",
        "waveLength": 324.726,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4399999999995998
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Fe",
        "waveLength": 324.721,
        "isSearchCriteriaMatched": true,
        "similarity": 0.33999999999969077
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 324.789,
        "isSearchCriteriaMatched": true,
        "similarity": 0.30000000000063665
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Co",
        "waveLength": 324.718,
        "isSearchCriteriaMatched": true,
        "similarity": 0.2799999999999727
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "V",
        "waveLength": 324.79,
        "isSearchCriteriaMatched": true,
        "similarity": 0.2799999999999727
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Fe",
        "waveLength": 324.717,
        "isSearchCriteriaMatched": true,
        "similarity": 0.2599999999993088
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 324.711,
        "isSearchCriteriaMatched": true,
        "similarity": 0.13999999999987267
      }
    ]
  },
  {
    "peak": {
      "x": 327.396,
      "y": 10168
    },
    "left": {
      "x": 326.795545454545,
      "y": 632.666666666667
    },
    "right": {
      "x": 328.488535714286,
      "y": 723
    },
    "area": 2768.1999334973566,
    "Elements": [
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 327.396,
        "isSearchCriteriaMatched": true,
        "similarity": 1
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Cu",
        "waveLength": 327.396,
        "isSearchCriteriaMatched": true,
        "similarity": 1
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ca",
        "waveLength": 327.395,
        "isSearchCriteriaMatched": true,
        "similarity": 0.9799999999993361
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Sb",
        "waveLength": 327.397,
        "isSearchCriteriaMatched": true,
        "similarity": 0.9800000000004729
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Co",
        "waveLength": 327.393,
        "isSearchCriteriaMatched": true,
        "similarity": 0.9399999999991451
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 327.392,
        "isSearchCriteriaMatched": true,
        "similarity": 0.919999999999618
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Th",
        "waveLength": 327.392,
        "isSearchCriteriaMatched": true,
        "similarity": 0.919999999999618
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "In",
        "waveLength": 327.402,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8800000000005639
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Nb",
        "waveLength": 327.389,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8599999999999
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ti",
        "waveLength": 327.404,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8400000000003729
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 327.406,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8000000000001819
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 327.411,
        "isSearchCriteriaMatched": true,
        "similarity": 0.7000000000002728
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Mo",
        "waveLength": 327.418,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5600000000001728
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Gd",
        "waveLength": 327.419,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5400000000006457
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Na",
        "waveLength": 327.422,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4799999999997908
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Hf",
        "waveLength": 327.365,
        "isSearchCriteriaMatched": true,
        "similarity": 0.37999999999988177
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Sc",
        "waveLength": 327.362,
        "isSearchCriteriaMatched": true,
        "similarity": 0.3200000000001637
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Mo",
        "waveLength": 327.357,
        "isSearchCriteriaMatched": true,
        "similarity": 0.22000000000025466
      },
      {
        "intensity": 3,
        "ionizationStage": 3,
        "element": "Fe",
        "waveLength": 327.353,
        "isSearchCriteriaMatched": true,
        "similarity": 0.13999999999987267
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "O",
        "waveLength": 327.352,
        "isSearchCriteriaMatched": true,
        "similarity": 0.11999999999920874
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ag",
        "waveLength": 327.441,
        "isSearchCriteriaMatched": true,
        "similarity": 0.10000000000081855
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Au",
        "waveLength": 327.351,
        "isSearchCriteriaMatched": true,
        "similarity": 0.09999999999968168
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 327.351,
        "isSearchCriteriaMatched": true,
        "similarity": 0.09999999999968168
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Nb",
        "waveLength": 327.35,
        "isSearchCriteriaMatched": true,
        "similarity": 0.08000000000015461
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Fe",
        "waveLength": 327.445,
        "isSearchCriteriaMatched": true,
        "similarity": 0.020000000000436557
      }
    ]
  },
  {
    "peak": {
      "x": 334.563882352941,
      "y": 8914.33333333333
    },
    "left": {
      "x": 333.546882352941,
      "y": 556
    },
    "right": {
      "x": 335.162117647059,
      "y": 574.666666666667
    },
    "area": 2208.789759337079,
    "Elements": [
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Fe",
        "waveLength": 334.568,
        "isSearchCriteriaMatched": true,
        "similarity": 0.9176470588201937
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Mo",
        "waveLength": 334.557,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8623529411804611
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Zn",
        "waveLength": 334.557,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8623529411804611
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Be",
        "waveLength": 334.543,
        "isSearchCriteriaMatched": true,
        "similarity": 0.582352941180261
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 334.543,
        "isSearchCriteriaMatched": true,
        "similarity": 0.582352941180261
      },
      {
        "intensity": 3,
        "ionizationStage": 3,
        "element": "W",
        "waveLength": 334.586,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5576470588196116
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "U",
        "waveLength": 334.589,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4976470588198936
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "V",
        "waveLength": 334.59,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4776470588203665
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Cr",
        "waveLength": 334.536,
        "isSearchCriteriaMatched": true,
        "similarity": 0.442352941180161
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Mn",
        "waveLength": 334.535,
        "isSearchCriteriaMatched": true,
        "similarity": 0.42235294118063393
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Ni",
        "waveLength": 334.593,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4176470588195116
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Zn",
        "waveLength": 334.593,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4176470588195116
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "In",
        "waveLength": 334.594,
        "isSearchCriteriaMatched": true,
        "similarity": 0.3976470588199845
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "K",
        "waveLength": 334.532,
        "isSearchCriteriaMatched": true,
        "similarity": 0.362352941179779
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ru",
        "waveLength": 334.532,
        "isSearchCriteriaMatched": true,
        "similarity": 0.362352941179779
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Cr",
        "waveLength": 334.601,
        "isSearchCriteriaMatched": true,
        "similarity": 0.2576470588198845
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Ni",
        "waveLength": 334.602,
        "isSearchCriteriaMatched": true,
        "similarity": 0.23764705882035742
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Er",
        "waveLength": 334.603,
        "isSearchCriteriaMatched": true,
        "similarity": 0.21764705881969348
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 334.523,
        "isSearchCriteriaMatched": true,
        "similarity": 0.18235294118062484
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "In",
        "waveLength": 334.606,
        "isSearchCriteriaMatched": true,
        "similarity": 0.15764705881997543
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "W",
        "waveLength": 334.611,
        "isSearchCriteriaMatched": true,
        "similarity": 0.05764705882006638
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Mo",
        "waveLength": 334.612,
        "isSearchCriteriaMatched": true,
        "similarity": 0.037647058819402446
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Cr",
        "waveLength": 334.514,
        "isSearchCriteriaMatched": true,
        "similarity": 0.0023529411803338007
      }
    ]
  },
  {
    "peak": {
      "x": 330.309428571429,
      "y": 6416.33333333333
    },
    "left": {
      "x": 329.884553571429,
      "y": 670.333333333333
    },
    "right": {
      "x": 331.034294117647,
      "y": 664.333333333333
    },
    "area": 1588.5561326854172,
    "Elements": [
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "La",
        "waveLength": 330.31,
        "isSearchCriteriaMatched": true,
        "similarity": 0.9885714285801441
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Mo",
        "waveLength": 330.311,
        "isSearchCriteriaMatched": true,
        "similarity": 0.9685714285806171
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Na",
        "waveLength": 330.298,
        "isSearchCriteriaMatched": true,
        "similarity": 0.7714285714198468
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Ce",
        "waveLength": 330.322,
        "isSearchCriteriaMatched": true,
        "similarity": 0.748571428580135
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ca",
        "waveLength": 330.294,
        "isSearchCriteriaMatched": true,
        "similarity": 0.6914285714194648
      },
      {
        "intensity": 3,
        "ionizationStage": 3,
        "element": "Zn",
        "waveLength": 330.294,
        "isSearchCriteriaMatched": true,
        "similarity": 0.6914285714194648
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 330.291,
        "isSearchCriteriaMatched": true,
        "similarity": 0.6314285714197467
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Mn",
        "waveLength": 330.328,
        "isSearchCriteriaMatched": true,
        "similarity": 0.6285714285806989
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Cr",
        "waveLength": 330.287,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5514285714193647
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Nb",
        "waveLength": 330.332,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5485714285803169
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "W",
        "waveLength": 330.333,
        "isSearchCriteriaMatched": true,
        "similarity": 0.528571428579653
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Fe",
        "waveLength": 330.285,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5114285714203106
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Mo",
        "waveLength": 330.334,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5085714285801259
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ta",
        "waveLength": 330.277,
        "isSearchCriteriaMatched": true,
        "similarity": 0.35142857141954664
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Fe",
        "waveLength": 330.347,
        "isSearchCriteriaMatched": true,
        "similarity": 0.24857142858058978
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Th",
        "waveLength": 330.348,
        "isSearchCriteriaMatched": true,
        "similarity": 0.22857142857992585
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Mo",
        "waveLength": 330.268,
        "isSearchCriteriaMatched": true,
        "similarity": 0.1714285714192556
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Cu",
        "waveLength": 330.351,
        "isSearchCriteriaMatched": true,
        "similarity": 0.1685714285802078
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Mo",
        "waveLength": 330.352,
        "isSearchCriteriaMatched": true,
        "similarity": 0.14857142858068073
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Zr",
        "waveLength": 330.266,
        "isSearchCriteriaMatched": true,
        "similarity": 0.13142857142020148
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Nb",
        "waveLength": 330.262,
        "isSearchCriteriaMatched": true,
        "similarity": 0.05142857141981949
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Fe",
        "waveLength": 330.357,
        "isSearchCriteriaMatched": true,
        "similarity": 0.04857142857963481
      }
    ]
  },
  {
    "peak": {
      "x": 328.24575,
      "y": 3173
    },
    "left": {
      "x": 326.795545454545,
      "y": 632.666666666667
    },
    "right": {
      "x": 328.488535714286,
      "y": 723
    },
    "area": 2768.1999334973566,
    "Elements": [
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Mo",
        "waveLength": 328.245,
        "isSearchCriteriaMatched": true,
        "similarity": 0.9850000000003547
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Fe",
        "waveLength": 328.244,
        "isSearchCriteriaMatched": true,
        "similarity": 0.9650000000008276
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ag",
        "waveLength": 328.253,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8550000000000182
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "V",
        "waveLength": 328.253,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8550000000000182
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Ti",
        "waveLength": 328.233,
        "isSearchCriteriaMatched": true,
        "similarity": 0.7450000000003456
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Zn",
        "waveLength": 328.233,
        "isSearchCriteriaMatched": true,
        "similarity": 0.7450000000003456
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 328.231,
        "isSearchCriteriaMatched": true,
        "similarity": 0.7050000000001546
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Gd",
        "waveLength": 328.231,
        "isSearchCriteriaMatched": true,
        "similarity": 0.7050000000001546
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Gd",
        "waveLength": 328.226,
        "isSearchCriteriaMatched": true,
        "similarity": 0.6050000000002456
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Co",
        "waveLength": 328.223,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5450000000005275
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ni",
        "waveLength": 328.27,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5150000000001
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Cu",
        "waveLength": 328.272,
        "isSearchCriteriaMatched": true,
        "similarity": 0.47499999999990905
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Zr",
        "waveLength": 328.273,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4549999999992451
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ni",
        "waveLength": 328.283,
        "isSearchCriteriaMatched": true,
        "similarity": 0.254999999999427
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Zr",
        "waveLength": 328.283,
        "isSearchCriteriaMatched": true,
        "similarity": 0.254999999999427
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Co",
        "waveLength": 328.204,
        "isSearchCriteriaMatched": true,
        "similarity": 0.16500000000041837
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Fe",
        "waveLength": 328.289,
        "isSearchCriteriaMatched": true,
        "similarity": 0.1349999999999909
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Mo",
        "waveLength": 328.289,
        "isSearchCriteriaMatched": true,
        "similarity": 0.1349999999999909
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "B",
        "waveLength": 328.201,
        "isSearchCriteriaMatched": true,
        "similarity": 0.10500000000070031
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Be",
        "waveLength": 328.291,
        "isSearchCriteriaMatched": true,
        "similarity": 0.09499999999979991
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 328.199,
        "isSearchCriteriaMatched": true,
        "similarity": 0.06500000000050932
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Au",
        "waveLength": 328.293,
        "isSearchCriteriaMatched": true,
        "similarity": 0.05499999999960892
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Pt",
        "waveLength": 328.197,
        "isSearchCriteriaMatched": true,
        "similarity": 0.025000000000318323
      }
    ]
  },
  {
    "peak": {
      "x": 352.533024,
      "y": 2411.66666666667
    },
    "left": {
      "x": 351.937784,
      "y": 549.333333333333
    },
    "right": {
      "x": 353.544932,
      "y": 481.666666666667
    },
    "area": 452.47408993394004,
    "Elements": [
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Nb",
        "waveLength": 352.523,
        "isSearchCriteriaMatched": true,
        "similarity": 0.7995200000002569
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ti",
        "waveLength": 352.516,
        "isSearchCriteriaMatched": true,
        "similarity": 0.6595200000001569
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Co",
        "waveLength": 352.508,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4995199999993929
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Tb",
        "waveLength": 352.56,
        "isSearchCriteriaMatched": true,
        "similarity": 0.46048000000018874
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ru",
        "waveLength": 352.564,
        "isSearchCriteriaMatched": true,
        "similarity": 0.38047999999980675
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Tm",
        "waveLength": 352.502,
        "isSearchCriteriaMatched": true,
        "similarity": 0.3795199999999568
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "U",
        "waveLength": 352.565,
        "isSearchCriteriaMatched": true,
        "similarity": 0.3604800000002797
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Rh",
        "waveLength": 352.566,
        "isSearchCriteriaMatched": true,
        "similarity": 0.3404800000007526
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Mo",
        "waveLength": 352.498,
        "isSearchCriteriaMatched": true,
        "similarity": 0.2995199999995748
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ba",
        "waveLength": 352.497,
        "isSearchCriteriaMatched": true,
        "similarity": 0.27952000000004773
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "W",
        "waveLength": 352.573,
        "isSearchCriteriaMatched": true,
        "similarity": 0.20048000000065258
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Dy",
        "waveLength": 352.493,
        "isSearchCriteriaMatched": true,
        "similarity": 0.19951999999966574
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Nb",
        "waveLength": 352.493,
        "isSearchCriteriaMatched": true,
        "similarity": 0.19951999999966574
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Dy",
        "waveLength": 352.575,
        "isSearchCriteriaMatched": true,
        "similarity": 0.1604800000004616
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "V",
        "waveLength": 352.577,
        "isSearchCriteriaMatched": true,
        "similarity": 0.12048000000027059
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Ti",
        "waveLength": 352.487,
        "isSearchCriteriaMatched": true,
        "similarity": 0.07952000000022963
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Zr",
        "waveLength": 352.581,
        "isSearchCriteriaMatched": true,
        "similarity": 0.040479999999888605
      }
    ]
  },
  {
    "peak": {
      "x": 341.503411764706,
      "y": 2069
    },
    "left": {
      "x": 341.204294117647,
      "y": 462
    },
    "right": {
      "x": 341.802529411765,
      "y": 480.333333333333
    },
    "area": 354.06871271224645,
    "Elements": [
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 341.506,
        "isSearchCriteriaMatched": true,
        "similarity": 0.9482352941208774
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Zr",
        "waveLength": 341.495,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8317647058797775
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Ho",
        "waveLength": 341.49,
        "isSearchCriteriaMatched": true,
        "similarity": 0.7317647058798684
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "V",
        "waveLength": 341.488,
        "isSearchCriteriaMatched": true,
        "similarity": 0.6917647058796774
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Dy",
        "waveLength": 341.482,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5717647058802413
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Mo",
        "waveLength": 341.527,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5282352941205772
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Ta",
        "waveLength": 341.527,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5282352941205772
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ni",
        "waveLength": 341.477,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4717647058791954
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 341.53,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4682352941208592
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 341.476,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4517647058796683
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Co",
        "waveLength": 341.474,
        "isSearchCriteriaMatched": true,
        "similarity": 0.41176470587947733
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ca",
        "waveLength": 341.538,
        "isSearchCriteriaMatched": true,
        "similarity": 0.3082352941200952
      },
      {
        "intensity": 3,
        "ionizationStage": 3,
        "element": "Zr",
        "waveLength": 341.466,
        "isSearchCriteriaMatched": true,
        "similarity": 0.2517647058798502
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "W",
        "waveLength": 341.541,
        "isSearchCriteriaMatched": true,
        "similarity": 0.24823529412037715
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ru",
        "waveLength": 341.464,
        "isSearchCriteriaMatched": true,
        "similarity": 0.21176470587965923
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ca",
        "waveLength": 341.46,
        "isSearchCriteriaMatched": true,
        "similarity": 0.13176470587927724
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 341.46,
        "isSearchCriteriaMatched": true,
        "similarity": 0.13176470587927724
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Co",
        "waveLength": 341.553,
        "isSearchCriteriaMatched": true,
        "similarity": 0.008235294120368053
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Fe",
        "waveLength": 341.553,
        "isSearchCriteriaMatched": true,
        "similarity": 0.008235294120368053
      }
    ]
  },
  {
    "peak": {
      "x": 361.986291666667,
      "y": 2042.33333333333
    },
    "left": {
      "x": 361.808666666667,
      "y": 483
    },
    "right": {
      "x": 362.282333333333,
      "y": 497.666666666667
    },
    "area": 312.11189111913336,
    "Elements": [
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 361.992,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8858333333396331
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Yb",
        "waveLength": 361.98,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8741666666603578
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ag",
        "waveLength": 361.977,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8141666666595029
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Fe",
        "waveLength": 361.977,
        "isSearchCriteriaMatched": true,
        "similarity": 0.8141666666595029
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Tm",
        "waveLength": 361.996,
        "isSearchCriteriaMatched": true,
        "similarity": 0.805833333340388
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Dy",
        "waveLength": 361.997,
        "isSearchCriteriaMatched": true,
        "similarity": 0.785833333339724
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Nb",
        "waveLength": 361.973,
        "isSearchCriteriaMatched": true,
        "similarity": 0.7341666666602578
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ti",
        "waveLength": 362.001,
        "isSearchCriteriaMatched": true,
        "similarity": 0.7058333333404789
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Tb",
        "waveLength": 361.971,
        "isSearchCriteriaMatched": true,
        "similarity": 0.6941666666600668
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Sn",
        "waveLength": 362.008,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5658333333403789
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "U",
        "waveLength": 362.008,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5658333333403789
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Dy",
        "waveLength": 362.016,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4058333333396149
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Nb",
        "waveLength": 361.951,
        "isSearchCriteriaMatched": true,
        "similarity": 0.29416666666043056
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Fe",
        "waveLength": 362.022,
        "isSearchCriteriaMatched": true,
        "similarity": 0.2858333333401788
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Au",
        "waveLength": 362.023,
        "isSearchCriteriaMatched": true,
        "similarity": 0.26583333333951487
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Cr",
        "waveLength": 361.946,
        "isSearchCriteriaMatched": true,
        "similarity": 0.19416666666052151
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ti",
        "waveLength": 361.946,
        "isSearchCriteriaMatched": true,
        "similarity": 0.19416666666052151
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Dy",
        "waveLength": 361.945,
        "isSearchCriteriaMatched": true,
        "similarity": 0.17416666665985758
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Os",
        "waveLength": 361.943,
        "isSearchCriteriaMatched": true,
        "similarity": 0.1341666666596666
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 362.032,
        "isSearchCriteriaMatched": true,
        "similarity": 0.0858333333403607
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 361.939,
        "isSearchCriteriaMatched": true,
        "similarity": 0.05416666666042147
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Fe",
        "waveLength": 361.939,
        "isSearchCriteriaMatched": true,
        "similarity": 0.05416666666042147
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ni",
        "waveLength": 361.939,
        "isSearchCriteriaMatched": true,
        "similarity": 0.05416666666042147
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Zr",
        "waveLength": 361.939,
        "isSearchCriteriaMatched": true,
        "similarity": 0.05416666666042147
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Cu",
        "waveLength": 362.035,
        "isSearchCriteriaMatched": true,
        "similarity": 0.02583333333950577
      }
    ]
  },
  {
    "peak": {
      "x": 351.58064,
      "y": 1855
    },
    "left": {
      "x": 350.9854,
      "y": 656.666666666667
    },
    "right": {
      "x": 351.818736,
      "y": 518.666666666667
    },
    "area": 398.98992332736634,
    "Elements": [
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Zr",
        "waveLength": 351.58,
        "isSearchCriteriaMatched": true,
        "similarity": 0.9871999999993477
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Ce",
        "waveLength": 351.577,
        "isSearchCriteriaMatched": true,
        "similarity": 0.9271999999996297
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 351.593,
        "isSearchCriteriaMatched": true,
        "similarity": 0.7527999999999793
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Mo",
        "waveLength": 351.568,
        "isSearchCriteriaMatched": true,
        "similarity": 0.7471999999993386
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "W",
        "waveLength": 351.596,
        "isSearchCriteriaMatched": true,
        "similarity": 0.6928000000002612
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 351.563,
        "isSearchCriteriaMatched": true,
        "similarity": 0.6471999999994296
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "V",
        "waveLength": 351.601,
        "isSearchCriteriaMatched": true,
        "similarity": 0.5928000000003522
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Be",
        "waveLength": 351.554,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4671999999991385
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 351.554,
        "isSearchCriteriaMatched": true,
        "similarity": 0.4671999999991385
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "P",
        "waveLength": 351.615,
        "isSearchCriteriaMatched": true,
        "similarity": 0.31280000000015207
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Tb",
        "waveLength": 351.616,
        "isSearchCriteriaMatched": true,
        "similarity": 0.292800000000625
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Nb",
        "waveLength": 351.619,
        "isSearchCriteriaMatched": true,
        "similarity": 0.23279999999977008
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "Nb",
        "waveLength": 351.542,
        "isSearchCriteriaMatched": true,
        "similarity": 0.22719999999912943
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 351.62,
        "isSearchCriteriaMatched": true,
        "similarity": 0.21280000000024302
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ni",
        "waveLength": 351.622,
        "isSearchCriteriaMatched": true,
        "similarity": 0.17280000000005202
      },
      {
        "intensity": 2,
        "ionizationStage": 2,
        "element": "W",
        "waveLength": 351.622,
        "isSearchCriteriaMatched": true,
        "similarity": 0.17280000000005202
      },
      {
        "intensity": 1,
        "ionizationStage": 1,
        "element": "Ce",
        "waveLength": 351.539,
        "isSearchCriteriaMatched": true,
        "similarity": 0.16719999999941138
      }
    ]
  }
]
`
