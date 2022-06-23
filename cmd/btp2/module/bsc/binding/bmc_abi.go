package binding

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

type BMCMessage struct {
	Next string
	Seq  *big.Int
	Msg  []byte
	Raw  types.Log
}

const _ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_svc\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_sn\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_code\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_errMsg\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_svcErrCode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_svcErrMsg\",\"type\":\"string\"}],\"name\":\"ErrorOnBTPError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_next\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_seq\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"}],\"name\":\"Message\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_network\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_bmcManagementAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBmcBtpAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"requestAddService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"bsh\",\"type\":\"address\"}],\"internalType\":\"structTypes.Request[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_prev\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_msg\",\"type\":\"string\"}],\"name\":\"handleRelayMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rlp\",\"type\":\"bytes\"}],\"name\":\"decodeBTPMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"src\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dst\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"svc\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"sn\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.BMCMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"}],\"name\":\"tryDecodeBMCService\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"serviceType\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.BMCService\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"}],\"name\":\"tryDecodeGatherFeeMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"fa\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"svcs\",\"type\":\"string[]\"}],\"internalType\":\"structTypes.GatherFeeMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_svc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_sn\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_link\",\"type\":\"string\"}],\"name\":\"getStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rxSeq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txSeq\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"heightMTA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsetMTA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.VerifierStats\",\"name\":\"verifier\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"msgCount\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.RelayStats[]\",\"name\":\"relays\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"relayIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rotateHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rotateTerm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delayLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAggregation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rxHeightSrc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rxHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockIntervalSrc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockIntervalDst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.LinkStats\",\"name\":\"_linkStats\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

func UnpackEventLog(data []byte) (*BMCMessage, error) {
	bmcABI, err := abi.JSON(strings.NewReader(string(_ABI)))
	if err != nil {
		return nil, err
	}
	var message BMCMessage
	err = bmcABI.UnpackIntoInterface(&message, "Message", data)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
