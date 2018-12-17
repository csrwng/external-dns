// Code generated by "stringer -type=icmpCode"; DO NOT EDIT.

package cmd

import "strconv"

const (
	_icmpCode_name_0 = "netUnreachablehostUnreachableprotocolUnreachableportUnreachablefragmentationNeededAndDoNotFragmentWasSetsourceRouteFaileddestinationNetworkUnknowndestinationHostUnknownsourceHostIsolatedcommunicationWithDestinationNetworkIsAdminstrativelyProhibitedcommunicationWithDestinationHostIsAdminstrativelyProhibiteddestinationNetworkUnreachableForTypeOfServicedestinationHostUnreachableForTypeOfServicecommunicationAdministrativelyProhibitedhostPrecedenceViolationprecedenceCutoffInEffect"
	_icmpCode_name_1 = "redirectDatagramForTheNetworkredirectDatagramForTheHostredirectDatagramForTheTypeOfServiceAndNetworkredirectSDatagramForTheTypeOfSerivceAndHost"
	_icmpCode_name_2 = "alternateAddressForHost"
	_icmpCode_name_3 = "timeToLiveExceededInTransitfragmentReassemblyTimeExceeded"
	_icmpCode_name_4 = "pointerIndicatesTheErrormissingARequiredOptionbadLength"
)

var (
	_icmpCode_index_0 = [...]uint16{0, 14, 29, 48, 63, 104, 121, 146, 168, 186, 248, 307, 352, 394, 433, 456, 480}
	_icmpCode_index_1 = [...]uint8{0, 29, 55, 100, 143}
	_icmpCode_index_3 = [...]uint8{0, 27, 57}
	_icmpCode_index_4 = [...]uint8{0, 24, 46, 55}
)

func (i icmpCode) String() string {
	switch {
	case 768 <= i && i <= 783:
		i -= 768
		return _icmpCode_name_0[_icmpCode_index_0[i]:_icmpCode_index_0[i+1]]
	case 1280 <= i && i <= 1283:
		i -= 1280
		return _icmpCode_name_1[_icmpCode_index_1[i]:_icmpCode_index_1[i+1]]
	case i == 1536:
		return _icmpCode_name_2
	case 2816 <= i && i <= 2817:
		i -= 2816
		return _icmpCode_name_3[_icmpCode_index_3[i]:_icmpCode_index_3[i+1]]
	case 3072 <= i && i <= 3074:
		i -= 3072
		return _icmpCode_name_4[_icmpCode_index_4[i]:_icmpCode_index_4[i+1]]
	default:
		return "icmpCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
