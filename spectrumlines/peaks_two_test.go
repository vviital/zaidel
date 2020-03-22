package spectrumlines

import (
	"encoding/json"

	"github.com/vviital/zaidel/peaks"
)

var peaksTwo []peaks.Peak
var spectrumLinesTwo []DeterminedPeak

func init() {
	json.Unmarshal([]byte(peaksTwoJSON), &peaksTwo)
	json.Unmarshal([]byte(spectrumLinesTwoJSON), &spectrumLinesTwo)
}

var peaksTwoJSON = `
[
  {
      "peak": {
        "x": 279.580521126761,
        "y": 4474
      },
      "left": {
        "x": 278.916647887324,
        "y": 700
      },
      "right": {
        "x": 280.66685915493,
        "y": 710.8
      },
      "area": 1767.5708836294334
  },
  {
      "peak": {
        "x": 317.942832684825,
        "y": 3475.8
      },
      "left": {
        "x": 317.584350194553,
        "y": 618.2
      },
      "right": {
        "x": 318.719544747082,
        "y": 606.4
      },
      "area": 853.3411466994119
  },
  {
      "peak": {
        "x": 280.304746478873,
        "y": 3108.4
      },
      "left": {
        "x": 278.916647887324,
        "y": 700
      },
      "right": {
        "x": 280.66685915493,
        "y": 710.8
      },
      "area": 1767.5708836294334
  },
  {
    "peak": {
      "x": 315.911431906615,
      "y": 2258
    },
    "left": {
      "x": 315.254214007782,
      "y": 585.4
    },
    "right": {
      "x": 316.508902723735,
      "y": 609.2
    },
    "area": 494.3221314208553
  },
  {
    "peak": {
      "x": 334.982647058824,
      "y": 2358
    },
    "left": {
      "x": 334.743352941177,
      "y": 617.4
    },
    "right": {
      "x": 335.162117647059,
      "y": 705.2
    },
    "area": 350.0978177022407
  },
  {
    "peak": {
      "x": 373.76875,
      "y": 1702.6
    },
    "left": {
      "x": 373.4135,
      "y": 664.8
    },
    "right": {
      "x": 374.301625,
      "y": 691
    },
    "area": 294.3888658259252
  },
  {
    "peak": {
      "x": 336.179117647059,
      "y": 1657.6
    },
    "left": {
      "x": 335.999647058824,
      "y": 677.4
    },
    "right": {
      "x": 336.358588235294,
      "y": 641
    },
    "area": 206.8279377629624
  },
  {
    "peak": {
      "x": 285.25361971831,
      "y": 1554.8
    },
    "left": {
      "x": 284.770802816901,
      "y": 671.4
    },
    "right": {
      "x": 285.676084507042,
      "y": 657.2
    },
    "area": 223.13839522890964
  },
  {
    "peak": {
      "x": 358.544948,
      "y": 1368.4
    },
    "left": {
      "x": 357.652088,
      "y": 636.8
    },
    "right": {
      "x": 359.675904,
      "y": 586.2
    },
    "area": 592.1005133025457
  },
  {
    "peak": {
      "x": 323.499311284047,
      "y": 1310.8
    },
    "left": {
      "x": 323.200575875486,
      "y": 601
    },
    "right": {
      "x": 324.276023346303,
      "y": 702.6
    },
    "area": 413.99581675610534
  },
  {
    "peak": {
      "x": 337.315764705882,
      "y": 1421.6
    },
    "left": {
      "x": 337.076470588235,
      "y": 647
    },
    "right": {
      "x": 337.495235294118,
      "y": 629
    },
    "area": 184.33450872480535
  },
  {
    "peak": {
      "x": 368.558416666667,
      "y": 1109
    },
    "left": {
      "x": 367.966333333333,
      "y": 498.6
    },
    "right": {
      "x": 369.1505,
      "y": 498.2
    },
    "area": 124.40705575263414
  },
  {
    "peak": {
      "x": 338.452411764706,
      "y": 1174.4
    },
    "left": {
      "x": 337.973823529412,
      "y": 578.6
    },
    "right": {
      "x": 338.990823529412,
      "y": 548.8
    },
    "area": 171.95620835543332
  },
  {
    "peak": {
      "x": 370.689916666667,
      "y": 1034.4
    },
    "left": {
      "x": 370.275458333333,
      "y": 520.6
    },
    "right": {
      "x": 371.104375,
      "y": 546.2
    },
    "area": 136.0014214522721
  },
  {
    "peak": {
      "x": 376.018666666667,
      "y": 1223.2
    },
    "left": {
      "x": 375.781833333333,
      "y": 776.2
    },
    "right": {
      "x": 376.433125,
      "y": 752.8
    },
    "area": 203.20983031911237
  },
  {
    "peak": {
      "x": 283.744816901408,
      "y": 1153.4
    },
    "left": {
      "x": 282.718830985915,
      "y": 663.4
    },
    "right": {
      "x": 284.469042253521,
      "y": 670.4
    },
    "area": 200.99134794613445
  },
  {
    "peak": {
      "x": 334.204941176471,
      "y": 1068.8
    },
    "left": {
      "x": 334.025470588235,
      "y": 644
    },
    "right": {
      "x": 334.384411764706,
      "y": 648.2
    },
    "area": 102.35120055885182
  },
  {
    "peak": {
      "x": 308.858104072398,
      "y": 1095.4
    },
    "left": {
      "x": 308.677805429864,
      "y": 646
    },
    "right": {
      "x": 309.518494163424,
      "y": 599
    },
    "area": 172.93590859203687
  },
  {
    "peak": {
      "x": 359.080664,
      "y": 992
    },
    "left": {
      "x": 357.652088,
      "y": 636.8
    },
    "right": {
      "x": 359.675904,
      "y": 586.2
    },
    "area": 592.1005133025457
  }
]
`

var spectrumLinesTwoJSON = `
[
  {
    "peak": {
      "x": 279.580521126761,
      "y": 4474
    },
    "left": {
      "x": 278.916647887324,
      "y": 700
    },
    "right": {
      "x": 280.66685915493,
      "y": 710.8
    },
    "area": 1767.5708836294334,
    "Elements": [
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 279.581,
        "matched": true,
        "similarity": 0.9904225352202047
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Cr",
        "waveLength": 279.582,
        "matched": true,
        "similarity": 0.9704225352206777
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 279.585,
        "matched": true,
        "similarity": 0.9104225352209596
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Fe",
        "waveLength": 279.576,
        "matched": true,
        "similarity": 0.9095774647798862
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Nb",
        "waveLength": 279.586,
        "matched": true,
        "similarity": 0.8904225352202957
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 279.594,
        "matched": true,
        "similarity": 0.7304225352206686
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 279.56,
        "matched": true,
        "similarity": 0.5895774647794951
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 279.555,
        "matched": true,
        "similarity": 0.4895774647795861
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 279.555,
        "matched": true,
        "similarity": 0.4895774647795861
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Au",
        "waveLength": 279.553,
        "matched": true,
        "similarity": 0.4495774647793951
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Mg",
        "waveLength": 279.553,
        "matched": true,
        "similarity": 0.4495774647793951
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Mn",
        "waveLength": 279.611,
        "matched": true,
        "similarity": 0.39042253522075043
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Hg",
        "waveLength": 279.613,
        "matched": true,
        "similarity": 0.35042253522055944
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 279.614,
        "matched": true,
        "similarity": 0.3304225352210324
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 279.62,
        "matched": true,
        "similarity": 0.2104225352204594
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 279.623,
        "matched": true,
        "similarity": 0.15042253522074134
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 279.538,
        "matched": true,
        "similarity": 0.14957746477966793
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cu",
        "waveLength": 279.533,
        "matched": true,
        "similarity": 0.049577464779758884
      }
    ]
  },
  {
    "peak": {
      "x": 317.942832684825,
      "y": 3475.8
    },
    "left": {
      "x": 317.584350194553,
      "y": 618.2
    },
    "right": {
      "x": 318.719544747082,
      "y": 606.4
    },
    "area": 853.3411466994119,
    "Elements": [
      {
        "intensity": 2,
        "stage": 2,
        "element": "Y",
        "waveLength": 317.942,
        "matched": true,
        "similarity": 0.9833463034997294
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "W",
        "waveLength": 317.944,
        "matched": true,
        "similarity": 0.9766536965000796
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "V",
        "waveLength": 317.941,
        "matched": true,
        "similarity": 0.9633463034990655
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cr",
        "waveLength": 317.945,
        "matched": true,
        "similarity": 0.9566536965005525
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Fe",
        "waveLength": 317.95,
        "matched": true,
        "similarity": 0.8566536965006435
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "B",
        "waveLength": 317.935,
        "matched": true,
        "similarity": 0.8433463034996294
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 317.934,
        "matched": true,
        "similarity": 0.8233463035001023
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ca",
        "waveLength": 317.933,
        "matched": true,
        "similarity": 0.8033463034994384
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Hg",
        "waveLength": 317.932,
        "matched": true,
        "similarity": 0.7833463034999113
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 317.932,
        "matched": true,
        "similarity": 0.7833463034999113
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ti",
        "waveLength": 317.929,
        "matched": true,
        "similarity": 0.7233463034990564
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Cr",
        "waveLength": 317.928,
        "matched": true,
        "similarity": 0.7033463034995293
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Er",
        "waveLength": 317.961,
        "matched": true,
        "similarity": 0.6366536965001615
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ag",
        "waveLength": 317.924,
        "matched": true,
        "similarity": 0.6233463034991473
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Nb",
        "waveLength": 317.924,
        "matched": true,
        "similarity": 0.6233463034991473
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ru",
        "waveLength": 317.924,
        "matched": true,
        "similarity": 0.6233463034991473
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ag",
        "waveLength": 317.914,
        "matched": true,
        "similarity": 0.4233463034993292
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Rh",
        "waveLength": 317.973,
        "matched": true,
        "similarity": 0.39665369650015236
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Zr",
        "waveLength": 317.974,
        "matched": true,
        "similarity": 0.3766536965006253
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 317.977,
        "matched": true,
        "similarity": 0.31665369650090724
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cu",
        "waveLength": 317.979,
        "matched": true,
        "similarity": 0.27665369650071625
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Na",
        "waveLength": 317.906,
        "matched": true,
        "similarity": 0.2633463034997021
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 317.906,
        "matched": true,
        "similarity": 0.2633463034997021
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Pt",
        "waveLength": 317.9,
        "matched": true,
        "similarity": 0.14334630349912914
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 317.897,
        "matched": true,
        "similarity": 0.08334630349941108
      }
    ]
  },
  {
    "peak": {
      "x": 280.304746478873,
      "y": 3108.4
    },
    "left": {
      "x": 278.916647887324,
      "y": 700
    },
    "right": {
      "x": 280.66685915493,
      "y": 710.8
    },
    "area": 1767.5708836294334,
    "Elements": [
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 280.304,
        "matched": true,
        "similarity": 0.9850704225389109
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Tm",
        "waveLength": 280.31,
        "matched": true,
        "similarity": 0.8949295774605162
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 280.312,
        "matched": true,
        "similarity": 0.8549295774603252
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 280.312,
        "matched": true,
        "similarity": 0.8549295774603252
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 280.313,
        "matched": true,
        "similarity": 0.8349295774607981
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ni",
        "waveLength": 280.314,
        "matched": true,
        "similarity": 0.8149295774601342
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 280.295,
        "matched": true,
        "similarity": 0.8050704225397567
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 280.317,
        "matched": true,
        "similarity": 0.7549295774604161
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "In",
        "waveLength": 280.318,
        "matched": true,
        "similarity": 0.734929577460889
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Er",
        "waveLength": 280.287,
        "matched": true,
        "similarity": 0.6450704225389927
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Pt",
        "waveLength": 280.324,
        "matched": true,
        "similarity": 0.6149295774603161
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Eu",
        "waveLength": 280.284,
        "matched": true,
        "similarity": 0.5850704225392747
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ru",
        "waveLength": 280.281,
        "matched": true,
        "similarity": 0.5250704225395566
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mn",
        "waveLength": 280.28,
        "matched": true,
        "similarity": 0.5050704225388927
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "V",
        "waveLength": 280.28,
        "matched": true,
        "similarity": 0.5050704225388927
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "W",
        "waveLength": 280.33,
        "matched": true,
        "similarity": 0.49492957746087995
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Hg",
        "waveLength": 280.276,
        "matched": true,
        "similarity": 0.42507042253964755
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cr",
        "waveLength": 280.335,
        "matched": true,
        "similarity": 0.3949295774609709
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Nb",
        "waveLength": 280.273,
        "matched": true,
        "similarity": 0.3650704225399295
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Bi",
        "waveLength": 280.27,
        "matched": true,
        "similarity": 0.30507042253907457
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 280.27,
        "matched": true,
        "similarity": 0.30507042253907457
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Mg",
        "waveLength": 280.27,
        "matched": true,
        "similarity": 0.30507042253907457
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mn",
        "waveLength": 280.265,
        "matched": true,
        "similarity": 0.20507042253916552
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Hg",
        "waveLength": 280.347,
        "matched": true,
        "similarity": 0.1549295774609618
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "V",
        "waveLength": 280.347,
        "matched": true,
        "similarity": 0.1549295774609618
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Bi",
        "waveLength": 280.348,
        "matched": true,
        "similarity": 0.13492957746029788
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Cu",
        "waveLength": 280.256,
        "matched": true,
        "similarity": 0.02507042253887448
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "U",
        "waveLength": 280.256,
        "matched": true,
        "similarity": 0.02507042253887448
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Er",
        "waveLength": 280.354,
        "matched": true,
        "similarity": 0.014929577460861765
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Nb",
        "waveLength": 280.354,
        "matched": true,
        "similarity": 0.014929577460861765
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ru",
        "waveLength": 280.354,
        "matched": true,
        "similarity": 0.014929577460861765
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Bi",
        "waveLength": 280.255,
        "matched": true,
        "similarity": 0.005070422539347419
      }
    ]
  },
  {
    "peak": {
      "x": 315.911431906615,
      "y": 2258
    },
    "left": {
      "x": 315.254214007782,
      "y": 585.4
    },
    "right": {
      "x": 316.508902723735,
      "y": 609.2
    },
    "area": 494.3221314208553,
    "Elements": [
      {
        "intensity": 2,
        "stage": 2,
        "element": "Zr",
        "waveLength": 315.911,
        "matched": true,
        "similarity": 0.9913618676998794
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cr",
        "waveLength": 315.91,
        "matched": true,
        "similarity": 0.9713618677003524
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Pt",
        "waveLength": 315.908,
        "matched": true,
        "similarity": 0.9313618677001614
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ir",
        "waveLength": 315.915,
        "matched": true,
        "similarity": 0.9286381322997386
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 315.907,
        "matched": true,
        "similarity": 0.9113618676994975
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 315.918,
        "matched": true,
        "similarity": 0.8686381323000205
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Mo",
        "waveLength": 315.92,
        "matched": true,
        "similarity": 0.8286381322998295
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 315.902,
        "matched": true,
        "similarity": 0.8113618676995884
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Mo",
        "waveLength": 315.894,
        "matched": true,
        "similarity": 0.6513618676999613
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ru",
        "waveLength": 315.89,
        "matched": true,
        "similarity": 0.5713618676995793
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Nb",
        "waveLength": 315.933,
        "matched": true,
        "similarity": 0.5686381323002934
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Mo",
        "waveLength": 315.934,
        "matched": true,
        "similarity": 0.5486381322996294
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 315.888,
        "matched": true,
        "similarity": 0.5313618676993883
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ca",
        "waveLength": 315.887,
        "matched": true,
        "similarity": 0.5113618676998613
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "V",
        "waveLength": 315.936,
        "matched": true,
        "similarity": 0.5086381323005753
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 315.937,
        "matched": true,
        "similarity": 0.4886381322999114
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 315.941,
        "matched": true,
        "similarity": 0.40863813230066626
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 315.881,
        "matched": true,
        "similarity": 0.39136186769928827
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Cr",
        "waveLength": 315.881,
        "matched": true,
        "similarity": 0.39136186769928827
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 315.88,
        "matched": true,
        "similarity": 0.3713618676997612
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 315.878,
        "matched": true,
        "similarity": 0.3313618676995702
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mn",
        "waveLength": 315.874,
        "matched": true,
        "similarity": 0.2513618677003251
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ni",
        "waveLength": 315.952,
        "matched": true,
        "similarity": 0.18863813230018422
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Cr",
        "waveLength": 315.958,
        "matched": true,
        "similarity": 0.06863813229961124
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Gd",
        "waveLength": 315.863,
        "matched": true,
        "similarity": 0.03136186769984306
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cu",
        "waveLength": 315.862,
        "matched": true,
        "similarity": 0.011361867700315997
      }
    ]
  },
  {
    "peak": {
      "x": 334.982647058824,
      "y": 2358
    },
    "left": {
      "x": 334.743352941177,
      "y": 617.4
    },
    "right": {
      "x": 335.162117647059,
      "y": 705.2
    },
    "area": 350.0978177022407,
    "Elements": [
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 334.973,
        "matched": true,
        "similarity": 0.8070588235202649
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 334.996,
        "matched": true,
        "similarity": 0.7329411764803808
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Tm",
        "waveLength": 334.999,
        "matched": true,
        "similarity": 0.6729411764795259
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Er",
        "waveLength": 335.006,
        "matched": true,
        "similarity": 0.5329411764805627
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Gd",
        "waveLength": 335.009,
        "matched": true,
        "similarity": 0.4729411764797078
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Au",
        "waveLength": 334.954,
        "matched": true,
        "similarity": 0.4270588235201558
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 334.953,
        "matched": true,
        "similarity": 0.40705882351949185
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 334.952,
        "matched": true,
        "similarity": 0.3870588235199648
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Nb",
        "waveLength": 334.952,
        "matched": true,
        "similarity": 0.3870588235199648
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cu",
        "waveLength": 334.946,
        "matched": true,
        "similarity": 0.2670588235205287
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Cs",
        "waveLength": 334.945,
        "matched": true,
        "similarity": 0.24705882351986475
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ca",
        "waveLength": 335.021,
        "matched": true,
        "similarity": 0.2329411764796987
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ru",
        "waveLength": 335.025,
        "matched": true,
        "similarity": 0.15294117648045358
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 334.94,
        "matched": true,
        "similarity": 0.1470588235199557
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ti",
        "waveLength": 334.94,
        "matched": true,
        "similarity": 0.1470588235199557
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Er",
        "waveLength": 335.026,
        "matched": true,
        "similarity": 0.13294117647978965
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 335.028,
        "matched": true,
        "similarity": 0.09294117647959865
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 335.029,
        "matched": true,
        "similarity": 0.07294117648007159
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Nb",
        "waveLength": 334.935,
        "matched": true,
        "similarity": 0.047058823520046644
      },
      {
        "intensity": 3,
        "stage": 3,
        "element": "Cr",
        "waveLength": 334.934,
        "matched": true,
        "similarity": 0.02705882352051958
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "W",
        "waveLength": 334.934,
        "matched": true,
        "similarity": 0.02705882352051958
      }
    ]
  },
  {
    "peak": {
      "x": 373.76875,
      "y": 1702.6
    },
    "left": {
      "x": 373.4135,
      "y": 664.8
    },
    "right": {
      "x": 374.301625,
      "y": 691
    },
    "area": 294.3888658259252,
    "Elements": [
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 373.773,
        "matched": true,
        "similarity": 0.9149999999997362
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cr",
        "waveLength": 373.755,
        "matched": true,
        "similarity": 0.7249999999996817
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 373.784,
        "matched": true,
        "similarity": 0.6950000000003911
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 373.751,
        "matched": true,
        "similarity": 0.6449999999992997
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Hf",
        "waveLength": 373.787,
        "matched": true,
        "similarity": 0.635000000000673
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 373.79,
        "matched": true,
        "similarity": 0.5749999999998181
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "V",
        "waveLength": 373.743,
        "matched": true,
        "similarity": 0.4849999999996726
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 373.795,
        "matched": true,
        "similarity": 0.47499999999990905
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ru",
        "waveLength": 373.741,
        "matched": true,
        "similarity": 0.4449999999994816
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Zr",
        "waveLength": 373.739,
        "matched": true,
        "similarity": 0.4049999999992906
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "V",
        "waveLength": 373.799,
        "matched": true,
        "similarity": 0.39500000000066393
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Al",
        "waveLength": 373.8,
        "matched": true,
        "similarity": 0.375
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "U",
        "waveLength": 373.804,
        "matched": true,
        "similarity": 0.2950000000007549
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Nd",
        "waveLength": 373.806,
        "matched": true,
        "similarity": 0.2550000000005639
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Rh",
        "waveLength": 373.727,
        "matched": true,
        "similarity": 0.1649999999992815
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 373.811,
        "matched": true,
        "similarity": 0.15500000000065484
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Zr",
        "waveLength": 373.811,
        "matched": true,
        "similarity": 0.15500000000065484
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Er",
        "waveLength": 373.816,
        "matched": true,
        "similarity": 0.055000000000745786
      }
    ]
  },
  {
    "peak": {
      "x": 336.179117647059,
      "y": 1657.6
    },
    "left": {
      "x": 335.999647058824,
      "y": 677.4
    },
    "right": {
      "x": 336.358588235294,
      "y": 641
    },
    "area": 206.8279377629624,
    "Elements": [
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cr",
        "waveLength": 336.177,
        "matched": true,
        "similarity": 0.9576470588203847
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ti",
        "waveLength": 336.182,
        "matched": true,
        "similarity": 0.9423529411797062
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 336.176,
        "matched": true,
        "similarity": 0.9376470588197208
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Th",
        "waveLength": 336.174,
        "matched": true,
        "similarity": 0.8976470588195298
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 336.185,
        "matched": true,
        "similarity": 0.8823529411799882
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 336.185,
        "matched": true,
        "similarity": 0.8823529411799882
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ni",
        "waveLength": 336.186,
        "matched": true,
        "similarity": 0.8623529411804611
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "C",
        "waveLength": 336.172,
        "matched": true,
        "similarity": 0.8576470588204756
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Er",
        "waveLength": 336.167,
        "matched": true,
        "similarity": 0.7576470588194297
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ca",
        "waveLength": 336.192,
        "matched": true,
        "similarity": 0.7423529411798881
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Sc",
        "waveLength": 336.194,
        "matched": true,
        "similarity": 0.7023529411796972
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ta",
        "waveLength": 336.164,
        "matched": true,
        "similarity": 0.6976470588197117
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 336.195,
        "matched": true,
        "similarity": 0.6823529411801701
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Th",
        "waveLength": 336.162,
        "matched": true,
        "similarity": 0.6576470588195207
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 336.197,
        "matched": true,
        "similarity": 0.6423529411799791
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ru",
        "waveLength": 336.2,
        "matched": true,
        "similarity": 0.582352941180261
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Y",
        "waveLength": 336.2,
        "matched": true,
        "similarity": 0.582352941180261
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Pb",
        "waveLength": 336.158,
        "matched": true,
        "similarity": 0.5776470588202756
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 336.156,
        "matched": true,
        "similarity": 0.5376470588200846
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ni",
        "waveLength": 336.156,
        "matched": true,
        "similarity": 0.5376470588200846
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 336.155,
        "matched": true,
        "similarity": 0.5176470588194206
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "V",
        "waveLength": 336.151,
        "matched": true,
        "similarity": 0.4376470588201755
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ti",
        "waveLength": 336.21,
        "matched": true,
        "similarity": 0.38235294118044294
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ca",
        "waveLength": 336.213,
        "matched": true,
        "similarity": 0.322352941179588
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Nb",
        "waveLength": 336.217,
        "matched": true,
        "similarity": 0.2423529411803429
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Rh",
        "waveLength": 336.218,
        "matched": true,
        "similarity": 0.22235294117967896
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 336.137,
        "matched": true,
        "similarity": 0.15764705881997543
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Cr",
        "waveLength": 336.222,
        "matched": true,
        "similarity": 0.14235294118043385
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Gd",
        "waveLength": 336.224,
        "matched": true,
        "similarity": 0.10235294118024285
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ca",
        "waveLength": 336.228,
        "matched": true,
        "similarity": 0.022352941179860863
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 336.228,
        "matched": true,
        "similarity": 0.022352941179860863
      }
    ]
  },
  {
    "peak": {
      "x": 285.25361971831,
      "y": 1554.8
    },
    "left": {
      "x": 284.770802816901,
      "y": 671.4
    },
    "right": {
      "x": 285.676084507042,
      "y": 657.2
    },
    "area": 223.13839522890964,
    "Elements": [
      {
        "intensity": 2,
        "stage": 2,
        "element": "Au",
        "waveLength": 285.254,
        "matched": true,
        "similarity": 0.9923943661997328
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ag",
        "waveLength": 285.253,
        "matched": true,
        "similarity": 0.9876056337996033
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "V",
        "waveLength": 285.253,
        "matched": true,
        "similarity": 0.9876056337996033
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Th",
        "waveLength": 285.25,
        "matched": true,
        "similarity": 0.9276056337998853
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Tm",
        "waveLength": 285.259,
        "matched": true,
        "similarity": 0.8923943661998237
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Hg",
        "waveLength": 285.242,
        "matched": true,
        "similarity": 0.7676056338002581
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ta",
        "waveLength": 285.236,
        "matched": true,
        "similarity": 0.6476056337996852
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 285.235,
        "matched": true,
        "similarity": 0.6276056338001581
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Na",
        "waveLength": 285.281,
        "matched": true,
        "similarity": 0.4523943661999965
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 285.223,
        "matched": true,
        "similarity": 0.387605633800149
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "V",
        "waveLength": 285.287,
        "matched": true,
        "similarity": 0.3323943662005604
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 285.29,
        "matched": true,
        "similarity": 0.2723943661997055
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ir",
        "waveLength": 285.213,
        "matched": true,
        "similarity": 0.1876056338003309
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mg",
        "waveLength": 285.213,
        "matched": true,
        "similarity": 0.1876056338003309
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 285.213,
        "matched": true,
        "similarity": 0.1876056338003309
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 285.295,
        "matched": true,
        "similarity": 0.17239436619979642
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "W",
        "waveLength": 285.21,
        "matched": true,
        "similarity": 0.12760563379947598
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Na",
        "waveLength": 285.301,
        "matched": true,
        "similarity": 0.05239436620036031
      }
    ]
  },
  {
    "peak": {
      "x": 358.544948,
      "y": 1368.4
    },
    "left": {
      "x": 357.652088,
      "y": 636.8
    },
    "right": {
      "x": 359.675904,
      "y": 586.2
    },
    "area": 592.1005133025457,
    "Elements": [
      {
        "intensity": 2,
        "stage": 2,
        "element": "Yb",
        "waveLength": 358.547,
        "matched": true,
        "similarity": 0.9589599999990241
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cr",
        "waveLength": 358.553,
        "matched": true,
        "similarity": 0.838959999999588
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Zr",
        "waveLength": 358.554,
        "matched": true,
        "similarity": 0.818960000000061
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 358.557,
        "matched": true,
        "similarity": 0.758959999999206
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 358.532,
        "matched": true,
        "similarity": 0.7410400000001118
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cr",
        "waveLength": 358.53,
        "matched": true,
        "similarity": 0.7010399999999208
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Te",
        "waveLength": 358.527,
        "matched": true,
        "similarity": 0.6410400000002028
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Mo",
        "waveLength": 358.567,
        "matched": true,
        "similarity": 0.558959999999388
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 358.571,
        "matched": true,
        "similarity": 0.47895999999900596
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 358.571,
        "matched": true,
        "similarity": 0.47895999999900596
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 358.516,
        "matched": true,
        "similarity": 0.4210400000008576
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Th",
        "waveLength": 358.577,
        "matched": true,
        "similarity": 0.35895999999956985
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Dy",
        "waveLength": 358.578,
        "matched": true,
        "similarity": 0.3389600000000428
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "C",
        "waveLength": 358.58,
        "matched": true,
        "similarity": 0.2989599999998518
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "C",
        "waveLength": 358.58,
        "matched": true,
        "similarity": 0.2989599999998518
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 358.581,
        "matched": true,
        "similarity": 0.27895999999918786
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Tb",
        "waveLength": 358.508,
        "matched": true,
        "similarity": 0.26104000000009364
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Dy",
        "waveLength": 358.506,
        "matched": true,
        "similarity": 0.22103999999990265
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "U",
        "waveLength": 358.584,
        "matched": true,
        "similarity": 0.2189599999994698
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Th",
        "waveLength": 358.505,
        "matched": true,
        "similarity": 0.20104000000037558
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ti",
        "waveLength": 358.585,
        "matched": true,
        "similarity": 0.19895999999994274
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "In",
        "waveLength": 358.586,
        "matched": true,
        "similarity": 0.1789599999992788
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Mo",
        "waveLength": 358.589,
        "matched": true,
        "similarity": 0.11895999999956075
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Th",
        "waveLength": 358.589,
        "matched": true,
        "similarity": 0.11895999999956075
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "C",
        "waveLength": 358.498,
        "matched": true,
        "similarity": 0.06104000000027554
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Nb",
        "waveLength": 358.497,
        "matched": true,
        "similarity": 0.04104000000074848
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 358.496,
        "matched": true,
        "similarity": 0.021040000000084547
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Gd",
        "waveLength": 358.496,
        "matched": true,
        "similarity": 0.021040000000084547
      }
    ]
  },
  {
    "peak": {
      "x": 323.499311284047,
      "y": 1310.8
    },
    "left": {
      "x": 323.200575875486,
      "y": 601
    },
    "right": {
      "x": 324.276023346303,
      "y": 702.6
    },
    "area": 413.99581675610534,
    "Elements": [
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ca",
        "waveLength": 323.499,
        "matched": true,
        "similarity": 0.9937743190605488
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mn",
        "waveLength": 323.5,
        "matched": true,
        "similarity": 0.9862256809399241
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 323.498,
        "matched": true,
        "similarity": 0.9737743190598849
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 323.501,
        "matched": true,
        "similarity": 0.9662256809403971
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Mo",
        "waveLength": 323.501,
        "matched": true,
        "similarity": 0.9662256809403971
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Na",
        "waveLength": 323.492,
        "matched": true,
        "similarity": 0.8537743190604488
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Co",
        "waveLength": 323.489,
        "matched": true,
        "similarity": 0.7937743190595938
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Zr",
        "waveLength": 323.48,
        "matched": true,
        "similarity": 0.6137743190604397
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Os",
        "waveLength": 323.473,
        "matched": true,
        "similarity": 0.4737743190603396
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "V",
        "waveLength": 323.472,
        "matched": true,
        "similarity": 0.4537743190596757
      },
      {
        "intensity": 3,
        "stage": 3,
        "element": "Ta",
        "waveLength": 323.469,
        "matched": true,
        "similarity": 0.39377431905995763
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mn",
        "waveLength": 323.531,
        "matched": true,
        "similarity": 0.3662256809398059
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ni",
        "waveLength": 323.465,
        "matched": true,
        "similarity": 0.31377431905957565
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 323.462,
        "matched": true,
        "similarity": 0.2537743190598576
      },
      {
        "intensity": 3,
        "stage": 3,
        "element": "Mo",
        "waveLength": 323.538,
        "matched": true,
        "similarity": 0.22622568093970585
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ti",
        "waveLength": 323.452,
        "matched": true,
        "similarity": 0.05377431906003949
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "V",
        "waveLength": 323.451,
        "matched": true,
        "similarity": 0.033774319060512425
      }
    ]
  },
  {
    "peak": {
      "x": 337.315764705882,
      "y": 1421.6
    },
    "left": {
      "x": 337.076470588235,
      "y": 647
    },
    "right": {
      "x": 337.495235294118,
      "y": 629
    },
    "area": 184.33450872480535,
    "Elements": [
      {
        "intensity": 1,
        "stage": 1,
        "element": "S",
        "waveLength": 337.319,
        "matched": true,
        "similarity": 0.9352941176391596
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Mo",
        "waveLength": 337.31,
        "matched": true,
        "similarity": 0.8847058823605494
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 337.323,
        "matched": true,
        "similarity": 0.8552941176399145
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 337.323,
        "matched": true,
        "similarity": 0.8552941176399145
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Pd",
        "waveLength": 337.3,
        "matched": true,
        "similarity": 0.6847058823607313
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Pt",
        "waveLength": 337.299,
        "matched": true,
        "similarity": 0.6647058823600673
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 337.292,
        "matched": true,
        "similarity": 0.5247058823599673
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Zr",
        "waveLength": 337.342,
        "matched": true,
        "similarity": 0.47529411763980534
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 337.345,
        "matched": true,
        "similarity": 0.4152941176389504
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 337.286,
        "matched": true,
        "similarity": 0.4047058823605312
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Pt",
        "waveLength": 337.279,
        "matched": true,
        "similarity": 0.26470588236043113
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ti",
        "waveLength": 337.279,
        "matched": true,
        "similarity": 0.26470588236043113
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ni",
        "waveLength": 337.353,
        "matched": true,
        "similarity": 0.2552941176393233
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Dy",
        "waveLength": 337.278,
        "matched": true,
        "similarity": 0.24470588236090407
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Er",
        "waveLength": 337.275,
        "matched": true,
        "similarity": 0.18470588236004915
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ti",
        "waveLength": 337.358,
        "matched": true,
        "similarity": 0.15529411763941425
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cu",
        "waveLength": 337.359,
        "matched": true,
        "similarity": 0.1352941176398872
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "K",
        "waveLength": 337.36,
        "matched": true,
        "similarity": 0.11529411763922326
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "P",
        "waveLength": 337.27,
        "matched": true,
        "similarity": 0.0847058823601401
      },
      {
        "intensity": 3,
        "stage": 3,
        "element": "Ca",
        "waveLength": 337.268,
        "matched": true,
        "similarity": 0.0447058823599491
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "V",
        "waveLength": 337.266,
        "matched": true,
        "similarity": 0.004705882360894975
      }
    ]
  },
  {
    "peak": {
      "x": 368.558416666667,
      "y": 1109
    },
    "left": {
      "x": 367.966333333333,
      "y": 498.6
    },
    "right": {
      "x": 369.1505,
      "y": 498.2
    },
    "area": 124.40705575263414,
    "Elements": [
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 368.559,
        "matched": true,
        "similarity": 0.9883333333389146
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Cr",
        "waveLength": 368.557,
        "matched": true,
        "similarity": 0.9716666666608944
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mn",
        "waveLength": 368.555,
        "matched": true,
        "similarity": 0.9316666666607034
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Zr",
        "waveLength": 368.562,
        "matched": true,
        "similarity": 0.9283333333391965
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 368.552,
        "matched": true,
        "similarity": 0.8716666666609854
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 368.577,
        "matched": true,
        "similarity": 0.6283333333394694
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Nd",
        "waveLength": 368.577,
        "matched": true,
        "similarity": 0.6283333333394694
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 368.539,
        "matched": true,
        "similarity": 0.6116666666603123
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Dy",
        "waveLength": 368.578,
        "matched": true,
        "similarity": 0.6083333333399423
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Cr",
        "waveLength": 368.525,
        "matched": true,
        "similarity": 0.33166666666011224
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 368.523,
        "matched": true,
        "similarity": 0.2916666666610581
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ti",
        "waveLength": 368.595,
        "matched": true,
        "similarity": 0.2683333333388873
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mn",
        "waveLength": 368.521,
        "matched": true,
        "similarity": 0.2516666666608671
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ti",
        "waveLength": 368.519,
        "matched": true,
        "similarity": 0.21166666666067613
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 368.6,
        "matched": true,
        "similarity": 0.16833333333897826
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 368.603,
        "matched": true,
        "similarity": 0.10833333333926021
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Nb",
        "waveLength": 368.512,
        "matched": true,
        "similarity": 0.07166666666057608
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Nd",
        "waveLength": 368.607,
        "matched": true,
        "similarity": 0.02833333333887822
      }
    ]
  },
  {
    "peak": {
      "x": 338.452411764706,
      "y": 1174.4
    },
    "left": {
      "x": 337.973823529412,
      "y": 578.6
    },
    "right": {
      "x": 338.990823529412,
      "y": 548.8
    },
    "area": 171.95620835543332,
    "Elements": [
      {
        "intensity": 1,
        "stage": 1,
        "element": "V",
        "waveLength": 338.46,
        "matched": true,
        "similarity": 0.8482352941209683
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 338.461,
        "matched": true,
        "similarity": 0.8282352941203044
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 338.462,
        "matched": true,
        "similarity": 0.8082352941207773
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Cr",
        "waveLength": 338.465,
        "matched": true,
        "similarity": 0.7482352941210593
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Nb",
        "waveLength": 338.465,
        "matched": true,
        "similarity": 0.7482352941210593
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Hf",
        "waveLength": 338.469,
        "matched": true,
        "similarity": 0.6682352941206773
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 338.433,
        "matched": true,
        "similarity": 0.6117647058792954
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Cr",
        "waveLength": 338.424,
        "matched": true,
        "similarity": 0.4317647058790044
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Cu",
        "waveLength": 338.481,
        "matched": true,
        "similarity": 0.4282352941206682
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "K",
        "waveLength": 338.486,
        "matched": true,
        "similarity": 0.32823529412075914
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "W",
        "waveLength": 338.489,
        "matched": true,
        "similarity": 0.2682352941210411
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Hf",
        "waveLength": 338.414,
        "matched": true,
        "similarity": 0.2317647058791863
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cu",
        "waveLength": 338.494,
        "matched": true,
        "similarity": 0.16823529411999516
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Dy",
        "waveLength": 338.409,
        "matched": true,
        "similarity": 0.13176470587927724
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Se",
        "waveLength": 338.498,
        "matched": true,
        "similarity": 0.08823529412075004
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 338.406,
        "matched": true,
        "similarity": 0.07176470587955919
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Tm",
        "waveLength": 338.5,
        "matched": true,
        "similarity": 0.04823529412055905
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Dy",
        "waveLength": 338.502,
        "matched": true,
        "similarity": 0.008235294120368053
      }
    ]
  },
  {
    "peak": {
      "x": 370.689916666667,
      "y": 1034.4
    },
    "left": {
      "x": 370.275458333333,
      "y": 520.6
    },
    "right": {
      "x": 371.104375,
      "y": 546.2
    },
    "area": 136.0014214522721,
    "Elements": [
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 370.693,
        "matched": true,
        "similarity": 0.9383333333400969
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 370.686,
        "matched": true,
        "similarity": 0.921666666659803
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Th",
        "waveLength": 370.677,
        "matched": true,
        "similarity": 0.7416666666606488
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 370.705,
        "matched": true,
        "similarity": 0.6983333333400878
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Nb",
        "waveLength": 370.708,
        "matched": true,
        "similarity": 0.6383333333392329
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mn",
        "waveLength": 370.665,
        "matched": true,
        "similarity": 0.5016666666606397
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Zr",
        "waveLength": 370.663,
        "matched": true,
        "similarity": 0.46166666666044875
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 370.717,
        "matched": true,
        "similarity": 0.45833333334007875
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Sm",
        "waveLength": 370.717,
        "matched": true,
        "similarity": 0.45833333334007875
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Au",
        "waveLength": 370.656,
        "matched": true,
        "similarity": 0.3216666666603487
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Os",
        "waveLength": 370.656,
        "matched": true,
        "similarity": 0.3216666666603487
      },
      {
        "intensity": 3,
        "stage": 3,
        "element": "O",
        "waveLength": 370.724,
        "matched": true,
        "similarity": 0.3183333333399787
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Pt",
        "waveLength": 370.653,
        "matched": true,
        "similarity": 0.26166666666063065
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Er",
        "waveLength": 370.652,
        "matched": true,
        "similarity": 0.24166666665996672
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "U",
        "waveLength": 370.729,
        "matched": true,
        "similarity": 0.21833333334006966
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 370.738,
        "matched": true,
        "similarity": 0.03833333333977862
      }
    ]
  },
  {
    "peak": {
      "x": 376.018666666667,
      "y": 1223.2
    },
    "left": {
      "x": 375.781833333333,
      "y": 776.2
    },
    "right": {
      "x": 376.433125,
      "y": 752.8
    },
    "area": 203.20983031911237,
    "Elements": [
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 376.017,
        "matched": true,
        "similarity": 0.9666666666598758
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "V",
        "waveLength": 376.024,
        "matched": true,
        "similarity": 0.8933333333400242
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 376.013,
        "matched": true,
        "similarity": 0.8866666666594938
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 376.005,
        "matched": true,
        "similarity": 0.7266666666598667
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ru",
        "waveLength": 376.002,
        "matched": true,
        "similarity": 0.6666666666601486
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "W",
        "waveLength": 376.038,
        "matched": true,
        "similarity": 0.6133333333398241
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 376.039,
        "matched": true,
        "similarity": 0.593333333340297
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 376.039,
        "matched": true,
        "similarity": 0.593333333340297
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Rh",
        "waveLength": 376.04,
        "matched": true,
        "similarity": 0.5733333333396331
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Gd",
        "waveLength": 376.047,
        "matched": true,
        "similarity": 0.43333333333953306
      },
      {
        "intensity": 3,
        "stage": 3,
        "element": "O",
        "waveLength": 375.987,
        "matched": true,
        "similarity": 0.36666666666042147
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 376.053,
        "matched": true,
        "similarity": 0.31333333334009694
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ru",
        "waveLength": 375.984,
        "matched": true,
        "similarity": 0.30666666665956654
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Nd",
        "waveLength": 375.979,
        "matched": true,
        "similarity": 0.2066666666596575
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 375.975,
        "matched": true,
        "similarity": 0.12666666666041237
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 376.063,
        "matched": true,
        "similarity": 0.11333333334027884
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Nb",
        "waveLength": 376.064,
        "matched": true,
        "similarity": 0.09333333333961491
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 375.969,
        "matched": true,
        "similarity": 0.006666666659839393
      }
    ]
  },
  {
    "peak": {
      "x": 283.744816901408,
      "y": 1153.4
    },
    "left": {
      "x": 282.718830985915,
      "y": 663.4
    },
    "right": {
      "x": 284.469042253521,
      "y": 670.4
    },
    "area": 200.99134794613445,
    "Elements": [
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cu",
        "waveLength": 283.737,
        "matched": true,
        "similarity": 0.8436619718404472
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 283.734,
        "matched": true,
        "similarity": 0.7836619718395923
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 283.732,
        "matched": true,
        "similarity": 0.7436619718405382
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Sb",
        "waveLength": 283.731,
        "matched": true,
        "similarity": 0.7236619718398742
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 283.759,
        "matched": true,
        "similarity": 0.7163380281597256
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Fe",
        "waveLength": 283.73,
        "matched": true,
        "similarity": 0.7036619718403472
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Th",
        "waveLength": 283.73,
        "matched": true,
        "similarity": 0.7036619718403472
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "C",
        "waveLength": 283.76,
        "matched": true,
        "similarity": 0.6963380281601985
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 283.728,
        "matched": true,
        "similarity": 0.6636619718401562
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Pt",
        "waveLength": 283.723,
        "matched": true,
        "similarity": 0.5636619718402471
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Zr",
        "waveLength": 283.723,
        "matched": true,
        "similarity": 0.5636619718402471
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Se",
        "waveLength": 283.721,
        "matched": true,
        "similarity": 0.5236619718400561
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "U",
        "waveLength": 283.719,
        "matched": true,
        "similarity": 0.48366197183986515
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 283.715,
        "matched": true,
        "similarity": 0.40366197183948316
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ag",
        "waveLength": 283.776,
        "matched": true,
        "similarity": 0.37633802815980744
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 283.776,
        "matched": true,
        "similarity": 0.37633802815980744
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Al",
        "waveLength": 283.786,
        "matched": true,
        "similarity": 0.17633802815998934
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 283.703,
        "matched": true,
        "similarity": 0.16366197183947406
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Au",
        "waveLength": 283.787,
        "matched": true,
        "similarity": 0.15633802816046227
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Cr",
        "waveLength": 283.787,
        "matched": true,
        "similarity": 0.15633802816046227
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 283.789,
        "matched": true,
        "similarity": 0.11633802816027128
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 283.79,
        "matched": true,
        "similarity": 0.09633802815960735
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 283.696,
        "matched": true,
        "similarity": 0.023661971840510887
      }
    ]
  },
  {
    "peak": {
      "x": 334.204941176471,
      "y": 1068.8
    },
    "left": {
      "x": 334.025470588235,
      "y": 644
    },
    "right": {
      "x": 334.384411764706,
      "y": 648.2
    },
    "area": 102.35120055885182,
    "Elements": [
      {
        "intensity": 1,
        "stage": 1,
        "element": "Cr",
        "waveLength": 334.202,
        "matched": true,
        "similarity": 0.9411764705803307
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Nb",
        "waveLength": 334.198,
        "matched": true,
        "similarity": 0.8611764705799487
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ti",
        "waveLength": 334.214,
        "matched": true,
        "similarity": 0.8188235294196602
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 334.195,
        "matched": true,
        "similarity": 0.8011764705802307
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 334.219,
        "matched": true,
        "similarity": 0.7188235294197511
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 334.19,
        "matched": true,
        "similarity": 0.7011764705803216
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "S",
        "waveLength": 334.188,
        "matched": true,
        "similarity": 0.6611764705801306
      },
      {
        "intensity": 3,
        "stage": 3,
        "element": "Ti",
        "waveLength": 334.188,
        "matched": true,
        "similarity": 0.6611764705801306
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 334.222,
        "matched": true,
        "similarity": 0.6588235294200331
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "La",
        "waveLength": 334.222,
        "matched": true,
        "similarity": 0.6588235294200331
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 334.186,
        "matched": true,
        "similarity": 0.6211764705799396
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Re",
        "waveLength": 334.225,
        "matched": true,
        "similarity": 0.5988235294191782
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Er",
        "waveLength": 334.184,
        "matched": true,
        "similarity": 0.5811764705808855
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 334.184,
        "matched": true,
        "similarity": 0.5811764705808855
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ca",
        "waveLength": 334.227,
        "matched": true,
        "similarity": 0.558823529420124
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "V",
        "waveLength": 334.228,
        "matched": true,
        "similarity": 0.5388235294194601
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 334.229,
        "matched": true,
        "similarity": 0.518823529419933
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 334.231,
        "matched": true,
        "similarity": 0.47882352941974204
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ru",
        "waveLength": 334.166,
        "matched": true,
        "similarity": 0.22117647058030343
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 334.164,
        "matched": true,
        "similarity": 0.18117647058011244
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "W",
        "waveLength": 334.246,
        "matched": true,
        "similarity": 0.1788235294200149
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Nb",
        "waveLength": 334.161,
        "matched": true,
        "similarity": 0.12117647058039438
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 334.253,
        "matched": true,
        "similarity": 0.03882352941991485
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ti",
        "waveLength": 334.155,
        "matched": true,
        "similarity": 0.0011764705798213981
      }
    ]
  },
  {
    "peak": {
      "x": 308.858104072398,
      "y": 1095.4
    },
    "left": {
      "x": 308.677805429864,
      "y": 646
    },
    "right": {
      "x": 309.518494163424,
      "y": 599
    },
    "area": 172.93590859203687,
    "Elements": [
      {
        "intensity": 1,
        "stage": 1,
        "element": "B",
        "waveLength": 308.86,
        "matched": true,
        "similarity": 0.9620814479601449
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Al",
        "waveLength": 308.852,
        "matched": true,
        "similarity": 0.8779185520390911
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Co",
        "waveLength": 308.868,
        "matched": true,
        "similarity": 0.8020814479605178
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Th",
        "waveLength": 308.847,
        "matched": true,
        "similarity": 0.777918552039182
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Re",
        "waveLength": 308.876,
        "matched": true,
        "similarity": 0.6420814479608907
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 308.879,
        "matched": true,
        "similarity": 0.5820814479600358
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 308.835,
        "matched": true,
        "similarity": 0.5379185520391729
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Os",
        "waveLength": 308.827,
        "matched": true,
        "similarity": 0.37791855203954583
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Nb",
        "waveLength": 308.826,
        "matched": true,
        "similarity": 0.35791855204001877
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Zr",
        "waveLength": 308.9,
        "matched": true,
        "similarity": 0.16208144796087254
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ca",
        "waveLength": 308.901,
        "matched": true,
        "similarity": 0.1420814479602086
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Cu",
        "waveLength": 308.813,
        "matched": true,
        "similarity": 0.09791855203934574
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "V",
        "waveLength": 308.813,
        "matched": true,
        "similarity": 0.09791855203934574
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 308.906,
        "matched": true,
        "similarity": 0.042081447960299556
      }
    ]
  },
  {
    "peak": {
      "x": 359.080664,
      "y": 992
    },
    "left": {
      "x": 357.652088,
      "y": 636.8
    },
    "right": {
      "x": 359.675904,
      "y": 586.2
    },
    "area": 592.1005133025457,
    "Elements": [
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 359.082,
        "matched": true,
        "similarity": 0.9732800000003863
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Er",
        "waveLength": 359.076,
        "matched": true,
        "similarity": 0.9067200000001776
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "C",
        "waveLength": 359.086,
        "matched": true,
        "similarity": 0.8932800000000043
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Mo",
        "waveLength": 359.073,
        "matched": true,
        "similarity": 0.8467199999993227
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Tm",
        "waveLength": 359.073,
        "matched": true,
        "similarity": 0.8467199999993227
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Nb",
        "waveLength": 359.09,
        "matched": true,
        "similarity": 0.8132800000007592
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Nb",
        "waveLength": 359.071,
        "matched": true,
        "similarity": 0.8067200000002686
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Hg",
        "waveLength": 359.095,
        "matched": true,
        "similarity": 0.7132799999997133
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Dy",
        "waveLength": 359.066,
        "matched": true,
        "similarity": 0.7067199999992226
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 359.1,
        "matched": true,
        "similarity": 0.6132799999998042
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Ce",
        "waveLength": 359.06,
        "matched": true,
        "similarity": 0.5867199999997865
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Fe",
        "waveLength": 359.059,
        "matched": true,
        "similarity": 0.5667200000002595
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Sc",
        "waveLength": 359.048,
        "matched": true,
        "similarity": 0.34671999999977743
      },
      {
        "intensity": 3,
        "stage": 3,
        "element": "Si",
        "waveLength": 359.047,
        "matched": true,
        "similarity": 0.32672000000025037
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Gd",
        "waveLength": 359.046,
        "matched": true,
        "similarity": 0.30671999999958643
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "W",
        "waveLength": 359.043,
        "matched": true,
        "similarity": 0.24671999999986838
      },
      {
        "intensity": 2,
        "stage": 2,
        "element": "Nb",
        "waveLength": 359.12,
        "matched": true,
        "similarity": 0.213280000000168
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ca",
        "waveLength": 359.126,
        "matched": true,
        "similarity": 0.09328000000073189
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "Ce",
        "waveLength": 359.035,
        "matched": true,
        "similarity": 0.08672000000024127
      },
      {
        "intensity": 1,
        "stage": 1,
        "element": "In",
        "waveLength": 359.035,
        "matched": true,
        "similarity": 0.08672000000024127
      }
    ]
  }
]
`
