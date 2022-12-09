package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the getMoneySpent function below.
 */
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	 */
	 var arr = make([]int, 0)
    var total int
    for i := 0; i < len(keyboards); i++ {
        for j := 0; j < len(drives); j++ {
            total = int(keyboards[i] + drives[j])
            if int(b) >= total {
                arr = append(arr, total)
            }
        }
    }

    sort.Ints(arr)
    if(len(arr)==0){return -1}

    return int32(arr[len(arr)-1])
}

func main() {
	var keyboards = []int32{3, 1}
	var drives = []int32{5, 2, 8}
	fmt.Println(getMoneySpent(keyboards, drives, 10))
	var keyboards1 = []int32{4}
	var drives1 = []int32{5}
	fmt.Println(getMoneySpent(keyboards1, drives1, 5))
	//var keyboards2 = []int32{572413, 359082, 441954, 808540, 346797, 388146, 746621, 796145, 893712, 305826, 756960, 72862, 332076, 932850, 744610, 394641, 152015, 792872, 130348, 279208, 345705, 646511, 807885, 873508, 246404, 41654, 63268, 830823, 131740, 571462, 858939, 607609, 633927, 523276, 253259, 874673, 555421, 43444, 527699, 261088, 655584, 875960, 233241, 752516, 864598, 17650, 309439, 517123, 694289, 529692, 16129, 400973, 372742, 297461, 351790, 681879, 543607, 207838, 72753, 799129, 790310, 803219, 217830, 415355, 677191, 283612, 803394, 283943, 562004, 482400, 741051, 719742, 662010, 358950, 355610, 610181, 809041, 382571, 294768, 304592, 571342, 24866, 745566, 3915, 379042, 227259, 302836, 406498, 459450, 851453, 818148, 798665, 930132, 267226, 545207, 244596, 661616, 227374, 903478, 339314, 98570, 498311, 232131, 438767, 709060, 218323, 823931, 130481, 857223, 12581, 415201, 153998, 89361, 705516, 437881, 243694, 230787, 468708, 291263, 591435, 772442, 887041, 75545, 889056, 827768, 745902, 427766, 35165, 596199, 575200, 89673, 786761, 337436, 17292, 89722, 549809, 961761, 857958, 807444, 474209, 459737, 4108, 223653, 11016, 933666, 416730, 344675, 155044, 161021, 664601, 563292, 708959, 943987, 177920, 3832, 888734, 238179, 442008, 938660, 142879, 171125, 764424, 113269, 790153, 570316, 639440, 862038, 853658, 477218, 737540, 994789, 882405, 679325, 433986, 74092, 747766, 993056, 999977, 407987, 725095, 287468, 114210, 558005, 518436, 304377, 587481, 86477, 174250, 1429, 963109, 449685, 175620, 837052, 177598, 572445, 915374, 999433, 465311, 286912, 63203, 280519, 598323, 159786, 601962, 164101, 64516, 285251, 972742, 793019, 544153, 905961, 559922, 975307, 520233, 232342, 610870, 881497, 264360, 328330, 287307, 112866, 511455, 870221, 806868, 636249, 49876, 164769, 603541, 429929, 806960, 287323, 156953, 736961, 15026, 453519, 312784, 989503, 588277, 54758, 138315, 355096, 426219, 482009, 552446, 503247, 599019, 316007, 274474, 729405, 400120, 874945, 778515, 938931, 941427, 805275, 862075, 850772, 564411, 550670, 410386, 826075, 333032, 303573, 124985, 398569, 106576, 316985, 688513, 467493, 259853, 285644, 254720, 961829, 998420, 312916, 187969, 98067, 148094, 526936, 98708, 57112, 579936, 110569, 628006, 219439, 365311, 464455, 535534, 433349, 788578, 799873, 818289, 739642, 958128, 426415, 824454, 333780}
	//var drives2 = []int32{564969, 316541, 179698, 774170, 99787, 52085, 54525, 12825, 666676, 795724, 807719, 477254, 508799, 130911, 54727, 235619, 535592, 932715, 935961, 859697, 597748, 709945, 809610, 482334, 800333, 520441, 932340, 975465, 770403, 965989, 747869, 215318, 234156, 811288, 59939, 197288, 517717, 213586, 124463, 676246, 962413, 405093, 289991, 931755, 853137, 89030, 387712, 511641, 323268, 918729, 935370, 904318, 363405, 156565, 390376, 366721, 106116, 619261, 670598, 159431, 288905, 968241, 585737, 914956, 591485, 166988, 719362, 121183, 745410, 85428, 514569, 909649, 331067, 193668, 837640, 416379, 309936, 448953, 507624, 599146, 401343, 71752, 855272, 226810, 432016, 931701, 386187, 908615, 689980, 781467, 602577, 19406, 107271, 327081, 508284, 628101, 735645, 275253, 992299, 909792, 629289, 634033, 370377, 71923, 226203, 909211, 873027, 957983, 814749, 77730, 233976, 532481, 200678, 303433, 715298, 392928, 315024, 179692, 306468, 578203, 152325, 958196, 264035, 503186, 427604, 82934, 733004, 620142, 336384, 340300, 328679, 856119, 284978, 434777, 478392, 192884, 180678, 191149, 485741, 768938, 704335, 265648, 29798, 767407, 9782, 205934, 75140, 321759, 535781, 465592, 854648, 943409, 4886, 931942, 936873, 206919, 106303, 358300, 842174, 149434, 482964, 510100, 659506, 460687, 510799, 900964, 477156, 930215, 835088, 383654, 756057, 491849, 483512, 896439, 111662, 806890, 115429, 3658, 801721, 320132, 529510, 421870, 481698, 229769, 680721, 199617, 859036, 37794, 629015, 643515, 763020, 262570, 765443, 52359, 356342, 231032, 274145, 802095, 801345, 851681, 377264, 483176, 903540, 416165, 69707, 978074, 58231, 200610, 927213, 215408, 648601, 291933, 379685, 172797, 755231, 379452, 672479, 911040, 531321, 615487, 98896, 412040, 812363, 934622, 866219, 48476, 172208, 210324, 850704, 566017, 878046, 615692, 549553, 39567, 656327, 403130, 382324, 649301, 281856, 698228, 722146, 638284, 862070, 320636, 654225, 93542, 511706, 427209, 158961, 128784, 53471, 947838, 445793, 475278, 49361, 317331, 569463, 766477, 382816, 498061, 757360,
	//	310104, 305718, 812151, 593117, 867348, 537805, 673160, 4714, 46976, 72406, 566451, 290434, 763467, 882434, 136534, 147049, 228723, 711055, 882694, 973586, 20127, 790746, 752897, 243001, 473881, 908570, 446873, 93129, 7099, 468248, 819583, 887742, 359906, 641411, 540888, 413738, 677454, 292879, 702788, 375184, 956196, 509057, 531460, 319116, 399810, 969187, 953919, 933288, 642255, 870868, 298218, 421925, 126493, 691458, 842745, 903033, 362264, 850622, 660619, 414421, 470923, 112412, 198975, 75648, 601823, 577195, 60697, 642773, 413436, 850035, 546739, 292581, 436146, 167509, 748841, 379092, 828112, 353655, 771417, 132499, 38822, 672435, 194173, 508807, 536729, 721286, 935033, 352771, 835145, 325776, 842725, 472683, 303277, 877605, 429545, 40291, 470812, 805015, 157408, 369071, 31076, 854812, 947117, 802384, 240767, 348840, 240003, 694033, 83398, 159669, 928717, 896484, 579868, 404583, 667188, 455729, 911783, 705919, 345699, 821026, 621648, 236396, 38809, 661500, 186815, 392087, 831434, 368786, 766942, 690182, 350722, 387271, 376896, 346187, 533675, 836077, 242, 101299, 241299, 676045, 208279, 248782, 250540, 15502, 806088, 166802, 279737, 606799, 191565, 709287, 504378, 21063, 865841, 36434, 241405, 902746, 78554, 73323, 356486, 894384, 747821, 228218, 550815, 740185, 816978, 309804, 959809, 858946, 792842, 700613, 603940, 401904, 493453, 76320, 40079, 732866, 984544, 617024, 383568, 2913, 152263, 9422, 666860, 808976, 284353, 259459, 585911, 636277, 586901,
	//	454317, 445831, 863625, 836307, 917168, 265027, 997033, 448734, 153719, 537814, 863159, 663775, 266082, 784149, 590949, 394897, 505887, 646193, 17993, 150645, 932982, 932859, 78345, 383146, 688374, 729686, 785184, 20346, 795468, 205234, 135526, 487193, 447671, 878451, 829616, 527374, 302748, 930080, 467855, 483906, 272856, 490560, 463656, 119774, 888038, 802838, 455227, 247022, 556038, 899816, 404068, 911030, 460221, 943523, 15178, 520348, 283810, 677247, 436681, 778692, 633940, 131045, 450154, 601093, 337894, 34306, 388187, 497679, 242348, 154878, 924703, 182965, 521861, 365130, 13413, 432535, 19451, 623313, 973181, 74446, 221501, 503806, 805790, 182667, 58289, 574105, 921371, 637549, 874614, 556835, 571222, 738462, 9001, 335685, 257869, 765279, 406915, 432865, 460168, 484294, 297746, 648249, 32956, 18363, 590261, 654142, 524273, 618725, 283065, 309917}
	//fmt.Println(getMoneySpent(keyboards2, drives2, 729580))
}
