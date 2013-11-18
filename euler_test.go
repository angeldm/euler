package euler

import (
	. "github.com/angeldm/euler/problem"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test(t *testing.T) {
	Convey("Solving Project Euler", t, func() {
		Convey("Problem 1", func() {
			So(P001(), ShouldEqual, 233168)
		})

		Convey("Problem 2", func() {
			So(P002(), ShouldEqual, 4613732)
		})

		Convey("Problem 3", func() {
			So(P003(), ShouldEqual, 6857)
		})

		Convey("Problem 4", func() {
			So(P004(), ShouldEqual, 906609)
		})

		Convey("Problem 5", func() {
			So(P005(), ShouldEqual, 232792560)
		})

		Convey("Problem 6", func() {
			So(P006(), ShouldEqual, 25164150)
		})

		Convey("Problem 7", func() {
			So(P007(), ShouldEqual, 104743)
		})
	})
}

// var Solutions = make([]string, 200)

// func init() {
// 	Solutions[1] = "233168"
// 	Solutions[2] = "4613732"
// 	Solutions[3] = "6857"
// 	Solutions[4] = "906609"
// 	Solutions[5] = "232792560"
// 	Solutions[6] = "25164150"
// 	Solutions[7] = "104743"
// 	Solutions[8] = "40824"
// 	Solutions[9] = "31875000"
// 	Solutions[10] = "142913828922"
// 	Solutions[11] = "70600674"
// 	Solutions[12] = "76576500"
// 	Solutions[13] = "5537376230"
// 	Solutions[14] = "837799"
// 	Solutions[15] = "137846528820"
// 	Solutions[16] = "1366"
// 	Solutions[17] = "21124"
// 	Solutions[18] = "1074"
// 	Solutions[19] = "171"
// 	Solutions[20] = "648"
// 	Solutions[21] = "31626"
// 	Solutions[22] = "871198282"
// 	Solutions[23] = "4179871"
// 	Solutions[24] = "2783915460"
// 	Solutions[25] = "4782"
// 	Solutions[26] = "983"
// 	Solutions[27] = "-59231"
// 	Solutions[28] = "669171001"
// 	Solutions[29] = "9183"
// 	Solutions[30] = "443839"
// 	Solutions[31] = "73682"
// 	Solutions[32] = "45228"
// 	Solutions[33] = "100"
// 	Solutions[34] = "40730"
// 	Solutions[35] = "55"
// 	Solutions[36] = "872187"
// 	Solutions[37] = "748317"
// 	Solutions[38] = "932718654"
// 	Solutions[39] = "840"
// 	Solutions[40] = "210"
// 	Solutions[41] = "7652413"
// 	Solutions[42] = "162"
// 	Solutions[43] = "16695334890"
// 	Solutions[44] = "5482660"
// 	Solutions[45] = "1533776805"
// 	Solutions[46] = "5777"
// 	Solutions[47] = "134043"
// 	Solutions[48] = "9110846700"
// 	Solutions[49] = "296962999629"
// 	Solutions[50] = "997651"
// 	Solutions[51] = "121313"
// 	Solutions[52] = "142857"
// 	Solutions[53] = "4075"
// 	Solutions[54] = "376"
// 	Solutions[55] = "249"
// 	Solutions[56] = "972"
// 	Solutions[57] = "153"
// 	Solutions[58] = "26241"
// 	Solutions[59] = "107359"
// 	Solutions[60] = "26033"
// 	Solutions[61] = "28684"
// 	Solutions[62] = "127035954683"
// 	Solutions[63] = "49"
// 	Solutions[64] = "1322"
// 	Solutions[65] = "272"
// 	Solutions[66] = "661"
// 	Solutions[67] = "7273"
// 	Solutions[68] = "6531031914842725"
// 	Solutions[69] = "510510"
// 	Solutions[70] = "8319823"
// 	Solutions[71] = "428570"
// 	Solutions[72] = "303963552391"
// 	Solutions[73] = "7295372"
// 	Solutions[74] = "402"
// 	Solutions[75] = "161667"
// 	Solutions[76] = "190569291"
// 	Solutions[77] = "71"
// 	Solutions[78] = "55374"
// 	Solutions[79] = "73162890"
// 	Solutions[80] = "40886"
// 	Solutions[81] = "427337"
// 	Solutions[82] = "260324"
// 	Solutions[83] = "425185"
// 	Solutions[84] = "101524"
// 	Solutions[85] = "2772"
// 	Solutions[86] = "1818"
// 	Solutions[87] = "1097343"
// 	Solutions[88] = "7587457"
// 	Solutions[89] = "743"
// 	Solutions[90] = "1217"
// 	Solutions[91] = "14234"
// 	Solutions[92] = "8581146"
// 	Solutions[93] = "1258"
// 	Solutions[94] = "518408346"
// 	Solutions[95] = "14316"
// 	Solutions[96] = "24702"
// 	Solutions[97] = "8739992577"
// 	Solutions[98] = "18769"
// 	Solutions[99] = "709"
// 	Solutions[100] = "756872327473"
// 	Solutions[101] = "37076114526"
// 	Solutions[102] = "228"
// 	Solutions[103] = "20313839404245"
// 	Solutions[104] = "329468"
// 	Solutions[105] = "73702"
// 	Solutions[106] = "21384"
// 	Solutions[107] = "259679"
// 	Solutions[108] = "180180"
// 	Solutions[109] = "38182"
// 	Solutions[110] = "9350130049860600"
// 	Solutions[111] = "612407567715"
// 	Solutions[112] = "1587000"
// 	Solutions[113] = "51161058134250"
// 	Solutions[114] = "16475640049"
// 	Solutions[115] = "168"
// 	Solutions[116] = "20492570929"
// 	Solutions[117] = "100808458960497"
// 	Solutions[118] = "44680"
// 	Solutions[119] = "248155780267521"
// 	Solutions[120] = "333082500"
// 	Solutions[121] = "2269"
// 	Solutions[122] = "1582"
// 	Solutions[123] = "21035"
// 	Solutions[124] = "21417"
// 	Solutions[125] = "2906969179"
// 	Solutions[126] = "18522"
// 	Solutions[127] = "18407904"
// 	Solutions[128] = "14516824220"
// 	Solutions[129] = "1000023"
// 	Solutions[130] = "149253"
// 	Solutions[131] = "173"
// 	Solutions[132] = "843296"
// 	Solutions[133] = "453647705"
// 	Solutions[134] = "18613426663617118"
// 	Solutions[135] = "4989"
// 	Solutions[136] = "2544559"
// 	Solutions[137] = "1120149658760"
// 	Solutions[138] = "1118049290473932"
// 	Solutions[139] = "10057761"
// 	Solutions[140] = "5673835352990"
// 	Solutions[141] = "878454337159"
// 	Solutions[142] = "1006193"
// 	Solutions[143] = "30758397"
// 	Solutions[144] = "354"
// 	Solutions[145] = "608720"
// 	Solutions[146] = "676333270"
// 	Solutions[147] = "846910284"
// 	Solutions[148] = "2129970655314432"
// 	Solutions[149] = "52852124"
// 	Solutions[150] = "-271248680"
// 	Solutions[151] = "464399"
// 	Solutions[152] = "301"
// 	Solutions[153] = "17971254122360635"
// 	Solutions[154] = "479742450"
// 	Solutions[155] = "3857447"
// 	Solutions[156] = "21295121502550"
// 	Solutions[157] = "53490"
// 	Solutions[158] = "409511334375"
// 	Solutions[159] = "14489159"
// 	Solutions[160] = "16576"
// 	Solutions[161] = "20574308184277971"
// 	Solutions[162] = "3D58725572C62302"
// 	Solutions[163] = "343047"
// 	Solutions[164] = "378158756814587"
// 	Solutions[165] = "2868868"
// 	Solutions[166] = "7130034"
// 	Solutions[167] = "3916160068885"
// 	Solutions[168] = "59206"
// 	Solutions[169] = "178653872807"
// 	Solutions[170] = "9857164023"
// 	Solutions[171] = "142989277"
// 	Solutions[172] = "227485267000992000"
// 	Solutions[173] = "1572729"
// 	Solutions[174] = "209566"
// 	Solutions[175] = "1,13717420,8"
// 	Solutions[176] = "96818198400000"
// 	Solutions[177] = "129325"
// 	Solutions[178] = "126461847755"
// 	Solutions[179] = "986262"
// 	Solutions[180] = "285196020571078987"
// 	Solutions[181] = "83735848679360680"
// 	Solutions[182] = "399788195976"
// 	Solutions[183] = "48861552"
// 	Solutions[184] = "1725323624056"
// 	Solutions[185] = "4640261571849533"
// 	Solutions[186] = "2325629"
// 	Solutions[187] = "17427258"
// }

// Solutions[188]="95962097"
// Solutions[189]="10834893628237824"
// Solutions[190]="371048281"
// Solutions[191]="1918080160"
// Solutions[192]="57060635927998347"
// Solutions[193]="684465067343069"
// Solutions[194]="61190912"
// Solutions[195]="75085391"
// Solutions[196]="322303240771079935"
// Solutions[197]="710637717"
// Solutions[198]="52374425"
// Solutions[199]="00396087"
// Solutions[200]="229161792008"
// Solutions[201]="115039000"
// Solutions[202]="1209002624"
// Solutions[203]="34029210557338"
// Solutions[204]="2944730"
// Solutions[205]="5731441"
// Solutions[206]="1389019170"
// Solutions[207]="44043947822"
// Solutions[208]="331951449665644800"
// Solutions[209]="15964587728784"
// Solutions[210]="1598174770174689458"
// Solutions[211]="1922364685"
// Solutions[212]="328968937309"
// Solutions[213]="721154"
// Solutions[214]="1677366278943"
// Solutions[215]="806844323190414"
// Solutions[216]="5437849"
// Solutions[217]="6273134"
// Solutions[218]="0"
// Solutions[219]="64564225042"
// Solutions[220]="139776,963904"
// Solutions[221]="1884161251122450"
// Solutions[222]="1590933"
// Solutions[223]="61614848"
// Solutions[224]="4137330"
// Solutions[225]="2009"
// Solutions[226]="11316017"
// Solutions[227]="618622"
// Solutions[228]="86226"
// Solutions[229]="11325263"
// Solutions[230]="850481152593119296"
// Solutions[231]="7526965179680"
// Solutions[232]=83648556"
// Solutions[233]="271204031455541309"
// Solutions[234]="1259187438574927161"
// Solutions[235]="002322108633"
// Solutions[236]="123/59"
// Solutions[237]="15836928"
// Solutions[238]="9922545104535661"
// Solutions[239]="001887854841"
// Solutions[240]="7448717393364181966"
// Solutions[241]="482316491800641154"
// Solutions[242]="997104142249036713"
// Solutions[243]="892371480"
// Solutions[244]="96356848"
// Solutions[245]="288084712410001"
// Solutions[246]="810834388"
// Solutions[247]="782252"
// Solutions[248]="23507044290"
// Solutions[249]="9275262564250418"
// Solutions[250]="1425480602091519"
// Solutions[251]="18946051"
// Solutions[252]="104924"
// Solutions[253]="492847"
// Solutions[254]="8184523820510"
// Solutions[255]="4474011180"
// Solutions[256]="85765680"
// Solutions[257]="139012411"
// Solutions[258]="12747994"
// Solutions[259]="20101196798"
// Solutions[260]="167542057"
// Solutions[261]="238890850232021"
// Solutions[262]=205"
// Solutions[263]="2039506520"
// Solutions[264]="2816417]="1055
// Solutions[265]="209110240768
// Solutions[266]="1096883702440585
// Solutions[267]="0]="999992836187
// Solutions[268]="785478606870985
// Solutions[269]="1311109198529286
// Solutions[270]="82282080
// Solutions[271]="4617456485273129588
// Solutions[272]="8495585919506151122
// Solutions[273]="2032447591196869022
// Solutions[274]="1601912348822
// Solutions[275]="15030564
// Solutions[276]="5777137137739632912
// Solutions[277]="1125977393124310
// Solutions[278]="1228215747273908452
// Solutions[279]="416577688
// Solutions[280]="430]="088247
// Solutions[281]="1485776387445623
// Solutions[282]="1098988351
// Solutions[283]="28038042525570324
// Solutions[284]="5a411d7b
// Solutions[285]="157055]="80999
// Solutions[286]="52]="6494571953
// Solutions[287]="313135496
// Solutions[288]="605857431263981935
// Solutions[289]="6567944538
// Solutions[290]="20444710234716473
// Solutions[291]="4037526
// Solutions[292]="3600060866
// Solutions[293]="2209
// Solutions[294]="789184709
// Solutions[295]="4884650818
// Solutions[296]="1137208419
// Solutions[297]="2252639041804718029
// Solutions[298]="1]="76882294
// Solutions[299]="549936643
// Solutions[300]="8]="0540771484375
// Solutions[301]="2178309
// Solutions[302]="1170060
// Solutions[303]="1111981904675169
// Solutions[304]="283988410192
// Solutions[305]="18174995535140
// Solutions[306]="852938
// Solutions[307]="0]="7311720251
// Solutions[308]="1539669807660924
// Solutions[309]="210139
// Solutions[310]="2586528661783
// Solutions[311]="2466018557
// Solutions[312]="324681947
// Solutions[313]="2057774861813004
// Solutions[314]="132]="52756426
// Solutions[315]="13625242
// Solutions[316]="542934735751917735
// Solutions[317]="1856532]="8455
// Solutions[318]="709313889
// Solutions[319]="268457129
// Solutions[320]="278157919195482643
// Solutions[321]="2470433131948040
// Solutions[322]="999998760323313995
// Solutions[323]="6]="3551758451
// Solutions[324]="96972774
// Solutions[325]="54672965
// Solutions[326]="1966666166408794329
// Solutions[327]="34315549139516
// Solutions[328]="260511850222
// Solutions[329]="199740353/29386561536000
// Solutions[330]="15955822
// Solutions[331]="467178235146843549
// Solutions[332]="2717]="751525
// Solutions[333]="3053105
// Solutions[334]="150320021261690835
// Solutions[335]="5032316
// Solutions[336]="CAGBIHEFJDK
// Solutions[337]="85068035
// Solutions[338]="15614292
// Solutions[339]="19823]="542204
// Solutions[340]="291504964
// Solutions[341]="56098610614277014
// Solutions[342]="5943040885644
// Solutions[343]="269533451410884183
// Solutions[344]="
// Solutions[345]="13938
// Solutions[346]="336108797689259276
// Solutions[347]="11109800204052
// Solutions[348]="1004195061
// Solutions[349]="115384615384614952
// Solutions[350]="84664213
// Solutions[351]="11762187201804552
// Solutions[352]="378563]="260589
// Solutions[353]="1]="2759860331
// Solutions[354]="
// Solutions[355]="1726545007
// Solutions[356]="28010159
// Solutions[357]="1739023853137
// Solutions[358]="3284144505
// Solutions[359]="40632119
// Solutions[360]="878825614395267072
// Solutions[361]="
// Solutions[362]="457895958010
// Solutions[363]="0]="0000372091
// Solutions[364]="44855254
// Solutions[365]="162619462356610313
// Solutions[366]="88351299
// Solutions[367]="
// Solutions[368]="253]="6135092068
// Solutions[369]="862400558448
// Solutions[370]="41791929448408
// Solutions[371]="40]="66368097
// Solutions[372]="301450082318807027
// Solutions[373]="727227472448913
// Solutions[374]="334420941
// Solutions[375]="7435327983715286168
// Solutions[376]="
// Solutions[377]="732385277
// Solutions[378]="147534623725724718
// Solutions[379]="132314136838185
// Solutions[380]="
// Solutions[381]="139602943319822
// Solutions[382]="
