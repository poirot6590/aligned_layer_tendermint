package keeper

import (
	"context"
	"encoding/base64"
	"strconv"

	kim "alignedlayer/verifiers/kimchi"
	"alignedlayer/x/verify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Kimchi(goCtx context.Context, msg *types.MsgKimchi) (*types.MsgKimchiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	result := verifykim(msg.Proof)
	event := sdk.NewEvent("verification_finished",
		sdk.NewAttribute("proof_verifies", strconv.FormatBool(result)),
		sdk.NewAttribute("prover", "KIMCHI"))

	ctx.EventManager().EmitEvent(event)

	return &types.MsgKimchiResponse{}, nil
}

func verifykim(proof string) bool {
	decodedBytes := make([]byte, kim.MAX_PROOF_SIZE)
	nDecoded, err := base64.StdEncoding.Decode(decodedBytes, []byte(proof))

	if err != nil {
		return false
	}
	vk := `3AAXxKwgAAAAAAAAAAUAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAJhlx2O8UpECKTRkNCEAAAAAAAAAAAAAAAAAAAA+zHibknZjo4+8t84+OaOuO+0zCFbRmvLa4fIWxg04wwyHoYmWXh9N8Lr7wF83zfzI7dNEuB8yzlxBD6lLpb8vFGdmZmb4eXhwcTCF0GRwT6eZmZmZmZmZmZmZmZmZmZkZzgABAAADAACXkZHEISb9MdGCS6ridDI8xDebvvtR9EXwrvamMLJK+/efNMkmAJGRxCEx9b4cf3pZ1Fbv8fuOCJV2QL1Z4fM1Nc5Yv01LUICvNYCRkcQh2ctnEP580WzltPh+qYANBtWKGCreC32VC4V5gi/I/RyAkZHEIXmeX9gGvEIOTmoAc8NheLeaZTSliIbRy/QOgIcHuiomAJGRxCFtQ1Ydd6zSfCPfcQb+4TZG952dqlwMgLWCI2wEnXVRDICRkcQh4x20CnckMylovyMXjGhnEgwDMEwv3fBMYR+FjxabgRgAkZHEIfDtpwKsfIIkrSeFc2YK9YKmLvVSqBNqlaBvlnABudgSgJ+RkcQhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAkZHEIQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQJGRxCEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAECRkcQhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAkZHEIQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQJGRxCEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAECRkcQhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAkZHEIQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQJGRxCEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAECRkcQhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAkZHEIQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQJGRxCEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAECRkcQhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAkZHEIQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQJGRxCEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAECRkcQhASImJlvOsuWox4vidXminDY2eHVj8rSqmckBYzhgIAmAkZHEIQEiJiZbzrLlqMeL4nV5opw2Nnh1Y/K0qpnJAWM4YCAJgJGRxCG9WwfReMHZymYzrK8sRMvL6h/6h+vaKkKFsf0PLuiIA4CRkcQhASImJlvOsuWox4vidXminDY2eHVj8rSqmckBYzhgIAmAkZHEIQEiJiZbzrLlqMeL4nV5opw2Nnh1Y/K0qpnJAWM4YCAJgJGRxCEBIiYmW86y5ajHi+J1eaKcNjZ4dWPytKqZyQFjOGAgCYDAwMDAwMCXxCABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMQg46IU6RM00Mrx64XfW9dSTXPV63qvdCp+stQL/cjNuQDEIFppZSb6MJxBLBDoZgTDwK0s2UQ92FuCMgNyEoHPvzMAxCARhI4sC/GLDp98jDTbRiG2x1r6oeMtD5YWTpUpu/SHAMQgk1i+nT7ya7BbuquowyYuC7ZdNkeWlNFe8WyfOzdx7ADEIBBtSRxyjCQM8EVkuFxYa0ByofFkH9NoTAhVTMyXkfMAxCCj7lmRIwwZA9LfxUMBPDVn4IOaKaTm0hLu25mXA924AMCS3AXqgaRDZWxskoGlSW5kZXioUG9zZWlkb26kQ3VycoGkQ2VsbJKBp1dpdG5lc3MGpEN1cnKBpENlbGySgatDb2VmZmljaWVudACkQ3VycoGoQ29uc3RhbnSBo01kc5IAAIGkQ2VsbJKBp1dpdG5lc3MApEN1cnKBo1BvdwelU3RvcmWjTXVso0FkZIGoQ29uc3RhbnSBo01kc5IAAYGkQ2VsbJKBp1dpdG5lc3MBpEN1cnKBo1BvdwelU3RvcmWjTXVso0FkZIGoQ29uc3RhbnSBo01kc5IAAoGkQ2VsbJKBp1dpdG5lc3MCpEN1cnKBo1BvdwelU3RvcmWjTXVso0FkZKNTdWKBqUNoYWxsZW5nZaVBbHBoYYGjUG93AYGkQ2VsbJKBp1dpdG5lc3MHpEN1cnKBpENlbGySgatDb2VmZmljaWVudAGkQ3VycoGoQ29uc3RhbnSBo01kc5IBAIGkTG9hZACjTXVso0FkZIGoQ29uc3RhbnSBo01kc5IBAYGkTG9hZAGjTXVso0FkZIGoQ29uc3RhbnSBo01kc5IBAoGkTG9hZAKjTXVso0FkZKNTdWKjTXVso0FkZIGpQ2hhbGxlbmdlpUFscGhhgaNQb3cCgaRDZWxskoGnV2l0bmVzcwikQ3VycoGkQ2VsbJKBq0NvZWZmaWNpZW50AqRDdXJygahDb25zdGFudIGjTWRzkgIAgaRMb2FkAKNNdWyjQWRkgahDb25zdGFudIGjTWRzkgIBgaRMb2FkAaNNdWyjQWRkgahDb25zdGFudIGjTWRzkgICgaRMb2FkAqNNdWyjQWRko1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdwOBpENlbGySgadXaXRuZXNzCaRDdXJygaRDZWxskoGrQ29lZmZpY2llbnQDpEN1cnKBqENvbnN0YW50gaNNZHOSAACBpENlbGySgadXaXRuZXNzBqRDdXJygaNQb3cHpVN0b3Jlo011bKNBZGSBqENvbnN0YW50gaNNZHOSAAGBpENlbGySgadXaXRuZXNzB6RDdXJygaNQb3cHpVN0b3Jlo011bKNBZGSBqENvbnN0YW50gaNNZHOSAAKBpENlbGySgadXaXRuZXNzCKRDdXJygaNQb3cHpVN0b3Jlo011bKNBZGSjU3Vio011bKNBZGSBqUNoYWxsZW5nZaVBbHBoYYGjUG93BIGkQ2VsbJKBp1dpdG5lc3MKpEN1cnKBpENlbGySgatDb2VmZmljaWVudASkQ3VycoGoQ29uc3RhbnSBo01kc5IBAIGkTG9hZAOjTXVso0FkZIGoQ29uc3RhbnSBo01kc5IBAYGkTG9hZASjTXVso0FkZIGoQ29uc3RhbnSBo01kc5IBAoGkTG9hZAWjTXVso0FkZKNTdWKjTXVso0FkZIGpQ2hhbGxlbmdlpUFscGhhgaNQb3cFgaRDZWxskoGnV2l0bmVzcwukQ3VycoGkQ2VsbJKBq0NvZWZmaWNpZW50BaRDdXJygahDb25zdGFudIGjTWRzkgIAgaRMb2FkA6NNdWyjQWRkgahDb25zdGFudIGjTWRzkgIBgaRMb2FkBKNNdWyjQWRkgahDb25zdGFudIGjTWRzkgICgaRMb2FkBaNNdWyjQWRko1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdwaBpENlbGySgadXaXRuZXNzDKRDdXJygaRDZWxskoGrQ29lZmZpY2llbnQGpEN1cnKBqENvbnN0YW50gaNNZHOSAACBpENlbGySgadXaXRuZXNzCaRDdXJygaNQb3cHpVN0b3Jlo011bKNBZGSBqENvbnN0YW50gaNNZHOSAAGBpENlbGySgadXaXRuZXNzCqRDdXJygaNQb3cHpVN0b3Jlo011bKNBZGSBqENvbnN0YW50gaNNZHOSAAKBpENlbGySgadXaXRuZXNzC6RDdXJygaNQb3cHpVN0b3Jlo011bKNBZGSjU3Vio011bKNBZGSBqUNoYWxsZW5nZaVBbHBoYYGjUG93B4GkQ2VsbJKBp1dpdG5lc3MNpEN1cnKBpENlbGySgatDb2VmZmljaWVudAekQ3VycoGoQ29uc3RhbnSBo01kc5IBAIGkTG9hZAajTXVso0FkZIGoQ29uc3RhbnSBo01kc5IBAYGkTG9hZAejTXVso0FkZIGoQ29uc3RhbnSBo01kc5IBAoGkTG9hZAijTXVso0FkZKNTdWKjTXVso0FkZIGpQ2hhbGxlbmdlpUFscGhhgaNQb3cIgaRDZWxskoGnV2l0bmVzcw6kQ3VycoGkQ2VsbJKBq0NvZWZmaWNpZW50CKRDdXJygahDb25zdGFudIGjTWRzkgIAgaRMb2FkBqNNdWyjQWRkgahDb25zdGFudIGjTWRzkgIBgaRMb2FkB6NNdWyjQWRkgahDb25zdGFudIGjTWRzkgICgaRMb2FkCKNNdWyjQWRko1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdwmBpENlbGySgadXaXRuZXNzA6RDdXJygaRDZWxskoGrQ29lZmZpY2llbnQJpEN1cnKBqENvbnN0YW50gaNNZHOSAACBpENlbGySgadXaXRuZXNzDKRDdXJygaNQb3cHpVN0b3Jlo011bKNBZGSBqENvbnN0YW50gaNNZHOSAAGBpENlbGySgadXaXRuZXNzDaRDdXJygaNQb3cHpVN0b3Jlo011bKNBZGSBqENvbnN0YW50gaNNZHOSAAKBpENlbGySgadXaXRuZXNzDqRDdXJygaNQb3cHpVN0b3Jlo011bKNBZGSjU3Vio011bKNBZGSBqUNoYWxsZW5nZaVBbHBoYYGjUG93CoGkQ2VsbJKBp1dpdG5lc3MEpEN1cnKBpENlbGySgatDb2VmZmljaWVudAqkQ3VycoGoQ29uc3RhbnSBo01kc5IBAIGkTG9hZAmjTXVso0FkZIGoQ29uc3RhbnSBo01kc5IBAYGkTG9hZAqjTXVso0FkZIGoQ29uc3RhbnSBo01kc5IBAoGkTG9hZAujTXVso0FkZKNTdWKjTXVso0FkZIGpQ2hhbGxlbmdlpUFscGhhgaNQb3cLgaRDZWxskoGnV2l0bmVzcwWkQ3VycoGkQ2VsbJKBq0NvZWZmaWNpZW50C6RDdXJygahDb25zdGFudIGjTWRzkgIAgaRMb2FkCaNNdWyjQWRkgahDb25zdGFudIGjTWRzkgIBgaRMb2FkCqNNdWyjQWRkgahDb25zdGFudIGjTWRzkgICgaRMb2FkC6NNdWyjQWRko1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdwyBpENlbGySgadXaXRuZXNzAKROZXh0gaRDZWxskoGrQ29lZmZpY2llbnQMpEN1cnKBqENvbnN0YW50gaNNZHOSAACBpENlbGySgadXaXRuZXNzA6RDdXJygaNQb3cHpVN0b3Jlo011bKNBZGSBqENvbnN0YW50gaNNZHOSAAGBpENlbGySgadXaXRuZXNzBKRDdXJygaNQb3cHpVN0b3Jlo011bKNBZGSBqENvbnN0YW50gaNNZHOSAAKBpENlbGySgadXaXRuZXNzBaRDdXJygaNQb3cHpVN0b3Jlo011bKNBZGSjU3Vio011bKNBZGSBqUNoYWxsZW5nZaVBbHBoYYGjUG93DYGkQ2VsbJKBp1dpdG5lc3MBpE5leHSBpENlbGySgatDb2VmZmljaWVudA2kQ3VycoGoQ29uc3RhbnSBo01kc5IBAIGkTG9hZAyjTXVso0FkZIGoQ29uc3RhbnSBo01kc5IBAYGkTG9hZA2jTXVso0FkZIGoQ29uc3RhbnSBo01kc5IBAoGkTG9hZA6jTXVso0FkZKNTdWKjTXVso0FkZIGpQ2hhbGxlbmdlpUFscGhhgaNQb3cOgaRDZWxskoGnV2l0bmVzcwKkTmV4dIGkQ2VsbJKBq0NvZWZmaWNpZW50DqRDdXJygahDb25zdGFudIGjTWRzkgIAgaRMb2FkDKNNdWyjQWRkgahDb25zdGFudIGjTWRzkgIBgaRMb2FkDaNNdWyjQWRkgahDb25zdGFudIGjTWRzkgICgaRMb2FkDqNNdWyjQWRko1N1YqNNdWyjQWRko011bIGkQ2VsbJKBpUluZGV4qlZhckJhc2VNdWykQ3VycoGkQ2VsbJKBp1dpdG5lc3MFpEN1cnKBpENlbGySgadXaXRuZXNzBqROZXh0gaRDZWxskoGnV2l0bmVzcwWkTmV4dIGkQ2VsbJKBp1dpdG5lc3MEpE5leHSBpENlbGySgadXaXRuZXNzA6ROZXh0gaRDZWxskoGnV2l0bmVzcwKkTmV4dIGkQ2VsbJKBp1dpdG5lc3MEpEN1cnKjRHVwo0FkZKNBZGSjRHVwo0FkZKNBZGSjRHVwo0FkZKNBZGSjRHVwo0FkZKNBZGSjRHVwo0FkZKNBZGSjU3VigalDaGFsbGVuZ2WlQWxwaGGBo1BvdwGBpENlbGySgadXaXRuZXNzAqROZXh0o0R1cKNNdWyBpENlbGySgadXaXRuZXNzAqROZXh0o1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdwKBpENlbGySgadXaXRuZXNzAqRDdXJygaRDZWxskoGnV2l0bmVzcwCkQ3VycqNTdWKBpENlbGySgadXaXRuZXNzB6ROZXh0o011bIGkQ2VsbJKBp1dpdG5lc3MDpEN1cnKBpENlbGySgadXaXRuZXNzAqROZXh0o0R1cKNBZGSBqENvbnN0YW50gadMaXRlcmFsxCABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKNTdWKBpENlbGySgadXaXRuZXNzAaRDdXJyo011bKNTdWKjU3Vio011bKNBZGSBqUNoYWxsZW5nZaVBbHBoYYGjUG93A4GkQ2VsbJKBp1dpdG5lc3MDpEN1cnKjRHVwo0FkZIGkQ2VsbJKBp1dpdG5lc3MCpEN1cnKBpENlbGySgadXaXRuZXNzB6ROZXh0gaRDZWxskoGnV2l0bmVzcwekTmV4dKNNdWylU3RvcmWBpENlbGySgadXaXRuZXNzAqRDdXJyo1N1YoGkQ2VsbJKBp1dpdG5lc3MApEN1cnKjU3Vio1N1YqVTdG9yZYGkQ2VsbJKBp1dpdG5lc3MHpE5leHSjTXVso1N1YqVTdG9yZYGkTG9hZBGjTXVsgaRMb2FkEIGkTG9hZBCjTXVsgaRDZWxskoGnV2l0bmVzcwekQ3VycoGkQ2VsbJKBp1dpdG5lc3MApEN1cnKjU3VigaRMb2FkD6NBZGSjTXVso1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdwSBpENlbGySgadXaXRuZXNzCKRDdXJygaRDZWxskoGnV2l0bmVzcwOkQ3VycqNBZGSBpExvYWQQo011bIGkQ2VsbJKBp1dpdG5lc3MCpEN1cnKBpENlbGySgadXaXRuZXNzB6RDdXJyo1N1YoGkTG9hZBGjTXVso1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdwWBpENlbGySgadXaXRuZXNzA6ROZXh0o0R1cKNNdWyBpENlbGySgadXaXRuZXNzA6ROZXh0o1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdwaBpENlbGySgadXaXRuZXNzB6RDdXJygaRDZWxskoGnV2l0bmVzcwCkQ3VycqNTdWKBpENlbGySgadXaXRuZXNzCKROZXh0o011bIGkQ2VsbJKBp1dpdG5lc3MIpEN1cnKBpENlbGySgadXaXRuZXNzA6ROZXh0o0R1cKNBZGSBqENvbnN0YW50gadMaXRlcmFsxCABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKNTdWKBpENlbGySgadXaXRuZXNzAaRDdXJyo011bKNTdWKjU3Vio011bKNBZGSBqUNoYWxsZW5nZaVBbHBoYYGjUG93B4GkQ2VsbJKBp1dpdG5lc3MIpEN1cnKjRHVwo0FkZIGkQ2VsbJKBp1dpdG5lc3MHpEN1cnKBpENlbGySgadXaXRuZXNzCKROZXh0gaRDZWxskoGnV2l0bmVzcwikTmV4dKNNdWylU3RvcmWBpENlbGySgadXaXRuZXNzB6RDdXJyo1N1YoGkQ2VsbJKBp1dpdG5lc3MApEN1cnKjU3Vio1N1YqVTdG9yZYGkQ2VsbJKBp1dpdG5lc3MIpE5leHSjTXVso1N1YqVTdG9yZYGkTG9hZBSjTXVsgaRMb2FkE4GkTG9hZBOjTXVsgaRDZWxskoGnV2l0bmVzcwmkQ3VycoGkQ2VsbJKBp1dpdG5lc3MApEN1cnKjU3VigaRMb2FkEqNBZGSjTXVso1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdwiBpENlbGySgadXaXRuZXNzCqRDdXJygaRDZWxskoGnV2l0bmVzcwikQ3VycqNBZGSBpExvYWQTo011bIGkQ2VsbJKBp1dpdG5lc3MHpEN1cnKBpENlbGySgadXaXRuZXNzCaRDdXJyo1N1YoGkTG9hZBSjTXVso1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdwmBpENlbGySgadXaXRuZXNzBKROZXh0o0R1cKNNdWyBpENlbGySgadXaXRuZXNzBKROZXh0o1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdwqBpENlbGySgadXaXRuZXNzCaRDdXJygaRDZWxskoGnV2l0bmVzcwCkQ3VycqNTdWKBpENlbGySgadXaXRuZXNzCaROZXh0o011bIGkQ2VsbJKBp1dpdG5lc3MKpEN1cnKBpENlbGySgadXaXRuZXNzBKROZXh0o0R1cKNBZGSBqENvbnN0YW50gadMaXRlcmFsxCABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKNTdWKBpENlbGySgadXaXRuZXNzAaRDdXJyo011bKNTdWKjU3Vio011bKNBZGSBqUNoYWxsZW5nZaVBbHBoYYGjUG93C4GkQ2VsbJKBp1dpdG5lc3MKpEN1cnKjRHVwo0FkZIGkQ2VsbJKBp1dpdG5lc3MJpEN1cnKBpENlbGySgadXaXRuZXNzCaROZXh0gaRDZWxskoGnV2l0bmVzcwmkTmV4dKNNdWylU3RvcmWBpENlbGySgadXaXRuZXNzCaRDdXJyo1N1YoGkQ2VsbJKBp1dpdG5lc3MApEN1cnKjU3Vio1N1YqVTdG9yZYGkQ2VsbJKBp1dpdG5lc3MJpE5leHSjTXVso1N1YqVTdG9yZYGkTG9hZBejTXVsgaRMb2FkFoGkTG9hZBajTXVsgaRDZWxskoGnV2l0bmVzcwukQ3VycoGkQ2VsbJKBp1dpdG5lc3MApEN1cnKjU3VigaRMb2FkFaNBZGSjTXVso1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdwyBpENlbGySgadXaXRuZXNzDKRDdXJygaRDZWxskoGnV2l0bmVzcwqkQ3VycqNBZGSBpExvYWQWo011bIGkQ2VsbJKBp1dpdG5lc3MJpEN1cnKBpENlbGySgadXaXRuZXNzC6RDdXJyo1N1YoGkTG9hZBejTXVso1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1Bvdw2BpENlbGySgadXaXRuZXNzBaROZXh0o0R1cKNNdWyBpENlbGySgadXaXRuZXNzBaROZXh0o1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1Bvdw6BpENlbGySgadXaXRuZXNzC6RDdXJygaRDZWxskoGnV2l0bmVzcwCkQ3VycqNTdWKBpENlbGySgadXaXRuZXNzCqROZXh0o011bIGkQ2VsbJKBp1dpdG5lc3MMpEN1cnKBpENlbGySgadXaXRuZXNzBaROZXh0o0R1cKNBZGSBqENvbnN0YW50gadMaXRlcmFsxCABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKNTdWKBpENlbGySgadXaXRuZXNzAaRDdXJyo011bKNTdWKjU3Vio011bKNBZGSBqUNoYWxsZW5nZaVBbHBoYYGjUG93D4GkQ2VsbJKBp1dpdG5lc3MMpEN1cnKjRHVwo0FkZIGkQ2VsbJKBp1dpdG5lc3MLpEN1cnKBpENlbGySgadXaXRuZXNzCqROZXh0gaRDZWxskoGnV2l0bmVzcwqkTmV4dKNNdWylU3RvcmWBpENlbGySgadXaXRuZXNzC6RDdXJyo1N1YoGkQ2VsbJKBp1dpdG5lc3MApEN1cnKjU3Vio1N1YqVTdG9yZYGkQ2VsbJKBp1dpdG5lc3MKpE5leHSjTXVso1N1YqVTdG9yZYGkTG9hZBqjTXVsgaRMb2FkGYGkTG9hZBmjTXVsgaRDZWxskoGnV2l0bmVzcw2kQ3VycoGkQ2VsbJKBp1dpdG5lc3MApEN1cnKjU3VigaRMb2FkGKNBZGSjTXVso1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdxCBpENlbGySgadXaXRuZXNzDqRDdXJygaRDZWxskoGnV2l0bmVzcwykQ3VycqNBZGSBpExvYWQZo011bIGkQ2VsbJKBp1dpdG5lc3MLpEN1cnKBpENlbGySgadXaXRuZXNzDaRDdXJyo1N1YoGkTG9hZBqjTXVso1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdxGBpENlbGySgadXaXRuZXNzBqROZXh0o0R1cKNNdWyBpENlbGySgadXaXRuZXNzBqROZXh0o1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdxKBpENlbGySgadXaXRuZXNzDaRDdXJygaRDZWxskoGnV2l0bmVzcwCkQ3VycqNTdWKBpENlbGySgadXaXRuZXNzC6ROZXh0o011bIGkQ2VsbJKBp1dpdG5lc3MOpEN1cnKBpENlbGySgadXaXRuZXNzBqROZXh0o0R1cKNBZGSBqENvbnN0YW50gadMaXRlcmFsxCABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKNTdWKBpENlbGySgadXaXRuZXNzAaRDdXJyo011bKNTdWKjU3Vio011bKNBZGSBqUNoYWxsZW5nZaVBbHBoYYGjUG93E4GkQ2VsbJKBp1dpdG5lc3MOpEN1cnKjRHVwo0FkZIGkQ2VsbJKBp1dpdG5lc3MNpEN1cnKBpENlbGySgadXaXRuZXNzC6ROZXh0gaRDZWxskoGnV2l0bmVzcwukTmV4dKNNdWylU3RvcmWBpENlbGySgadXaXRuZXNzDaRDdXJyo1N1YoGkQ2VsbJKBp1dpdG5lc3MApEN1cnKjU3Vio1N1YqVTdG9yZYGkQ2VsbJKBp1dpdG5lc3MLpE5leHSjTXVso1N1YqVTdG9yZYGkTG9hZB2jTXVsgaRMb2FkHIGkTG9hZByjTXVsgaRDZWxskoGnV2l0bmVzcwCkTmV4dIGkQ2VsbJKBp1dpdG5lc3MApEN1cnKjU3VigaRMb2FkG6NBZGSjTXVso1N1YqNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdxSBpENlbGySgadXaXRuZXNzAaROZXh0gaRDZWxskoGnV2l0bmVzcw6kQ3VycqNBZGSBpExvYWQco011bIGkQ2VsbJKBp1dpdG5lc3MNpEN1cnKBpENlbGySgadXaXRuZXNzAKROZXh0o1N1YoGkTG9hZB2jTXVso1N1YqNNdWyjQWRko011bKNBZGSBpENlbGySgaVJbmRleKtDb21wbGV0ZUFkZKRDdXJygaRDZWxskoGnV2l0bmVzcwqkQ3VycoGkQ2VsbJKBp1dpdG5lc3MCpEN1cnKBpENlbGySgadXaXRuZXNzAKRDdXJyo1N1YqVTdG9yZaNNdWyBqENvbnN0YW50gadMaXRlcmFsxCABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIGkQ2VsbJKBp1dpdG5lc3MHpEN1cnKjU3Vio1N1YoGpQ2hhbGxlbmdlpUFscGhhgaNQb3cBgaRDZWxskoGnV2l0bmVzcwekQ3VycoGkTG9hZB6jTXVso011bKNBZGSBqUNoYWxsZW5nZaVBbHBoYYGjUG93AoGkQ2VsbJKBp1dpdG5lc3MHpEN1cnKBpENlbGySgadXaXRuZXNzCKRDdXJyo0R1cKNBZGSBpENlbGySgadXaXRuZXNzAaRDdXJyo011bIGkQ2VsbJKBp1dpdG5lc3MApEN1cnKBpENlbGySgadXaXRuZXNzAKRDdXJyo011bKVTdG9yZaNEdXCjQWRko1N1YoGkTG9hZB+jU3Vio011bIGoQ29uc3RhbnSBp0xpdGVyYWzEIAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgaRDZWxskoGnV2l0bmVzcwekQ3VycqNTdWKBpExvYWQegaRDZWxskoGnV2l0bmVzcwikQ3VycqNNdWyBpENlbGySgadXaXRuZXNzA6RDdXJygaRDZWxskoGnV2l0bmVzcwGkQ3VycqNTdWKlU3RvcmWjU3Vio011bKNBZGSjTXVso0FkZIGpQ2hhbGxlbmdlpUFscGhhgaNQb3cDgaRDZWxskoGnV2l0bmVzcwCkQ3VycoGkQ2VsbJKBp1dpdG5lc3MCpEN1cnKjQWRkgaRDZWxskoGnV2l0bmVzcwSkQ3VycqNBZGSBpENlbGySgadXaXRuZXNzCKRDdXJygaRDZWxskoGnV2l0bmVzcwikQ3VycqNNdWyjU3Vio011bKNBZGSBqUNoYWxsZW5nZaVBbHBoYYGjUG93BIGkQ2VsbJKBp1dpdG5lc3MIpEN1cnKBpENlbGySgadXaXRuZXNzAKRDdXJygaRDZWxskoGnV2l0bmVzcwSkQ3VycqNTdWKjTXVsgaRDZWxskoGnV2l0bmVzcwGkQ3VycqNTdWKBpENlbGySgadXaXRuZXNzBaRDdXJyo1N1YqNNdWyjQWRk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bMQgCwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACjQWRkgaRDZWxskoGnV2l0bmVzcwykQ3VycqNNdWyBqENvbnN0YW50gadMaXRlcmFsxCD7////7DAtmRv5TAn8mEYiAAAAAAAAAAAAAAAAAAAAQKNBZGSBpENlbGySgadXaXRuZXNzDKRDdXJyo011bKNNdWyjQWRkgalDaGFsbGVuZ2WlQWxwaGGBo1BvdwqBpENlbGySgadXaXRuZXNzDaRDdXJygahDb25zdGFudIGnTGl0ZXJhbMQg+////+wwLZkb+UwJ/JhGIgAAAAAAAAAAAAAAAAAAAECjQWRkgaRDZWxskoGnV2l0bmVzcw2kQ3VycqNNdWyBqENvbnN0YW50gadMaXRlcmFsxCALAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKNBZGSBpENlbGySgadXaXRuZXNzDaRDdXJyo011bIGoQ29uc3RhbnSBp0xpdGVyYWzEIPv////sMC2ZG/lMCfyYRiIAAAAAAAAAAAAAAAAAAABAo0FkZIGkQ2VsbJKBp1dpdG5lc3MNpEN1cnKjTXVso011bKNBZGSjTXVso0FkZIGkQ2VsbJKBpUluZGV4p0dlbmVyaWOkQ3VycoGkQ2VsbJKBq0NvZWZmaWNpZW50AKRDdXJygaRDZWxskoGnV2l0bmVzcwCkQ3VycqNNdWyBpENlbGySgatDb2VmZmljaWVudAGkQ3VycoGkQ2VsbJKBp1dpdG5lc3MBpEN1cnKjTXVso0FkZIGkQ2VsbJKBq0NvZWZmaWNpZW50AqRDdXJygaRDZWxskoGnV2l0bmVzcwKkQ3VycqNNdWyjQWRkgaRDZWxskoGrQ29lZmZpY2llbnQDpEN1cnKBpENlbGySgadXaXRuZXNzAKRDdXJyo011bIGkQ2VsbJKBp1dpdG5lc3MBpEN1cnKjTXVso0FkZIGkQ2VsbJKBq0NvZWZmaWNpZW50BKRDdXJyo0FkZIGpQ2hhbGxlbmdlpUFscGhhgaNQb3cBgaRDZWxskoGrQ29lZmZpY2llbnQFpEN1cnKBpENlbGySgadXaXRuZXNzA6RDdXJyo011bIGkQ2VsbJKBq0NvZWZmaWNpZW50BqRDdXJygaRDZWxskoGnV2l0bmVzcwSkQ3VycqNNdWyjQWRkgaRDZWxskoGrQ29lZmZpY2llbnQHpEN1cnKBpENlbGySgadXaXRuZXNzBaRDdXJyo011bKNBZGSBpENlbGySgatDb2VmZmljaWVudAikQ3VycoGkQ2VsbJKBp1dpdG5lc3MDpEN1cnKjTXVsgaRDZWxskoGnV2l0bmVzcwSkQ3VycqNNdWyjQWRkgaRDZWxskoGrQ29lZmZpY2llbnQJpEN1cnKjQWRko011bKNBZGSjTXVso0FkZJCTGIKrUGVybXV0YXRpb26SFQOBpEdhdGWkWmVyb5IAFcQBAA==`

	pubInputBuffer := make([]byte, kim.MAX_PUB_INPUT_SIZE)
	pubInputLen, err := base64.StdEncoding.Decode(pubInputBuffer, []byte(vk))
	if err != nil {
		return false
	}

	return kim.VerifyKimchiProof(([kim.MAX_PROOF_SIZE]byte)(decodedBytes), uint(nDecoded), ([kim.MAX_PUB_INPUT_SIZE]byte)(pubInputBuffer), uint(pubInputLen))
}