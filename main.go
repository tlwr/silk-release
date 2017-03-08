package main

import (
	"encoding/json"

	"github.com/containernetworking/cni/pkg/ipam"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/types/current"
	"github.com/containernetworking/cni/pkg/version"
)

type NetConf struct {
	types.NetConf
}

func main() {
	cmdAdd := func(args *skel.CmdArgs) error {
		var netConf NetConf
		err := json.Unmarshal(args.StdinData, &netConf)
		if err != nil {
			panic(err)
		}
		result, err := ipam.ExecAdd(netConf.IPAM.Type, args.StdinData)
		if err != nil {
			panic(err)
		}

		cniResult, err := current.NewResultFromResult(result)
		if err != nil {
			panic(err)
		}

		cniResult.IPs[0].Interface = -1

		return types.PrintResult(cniResult, netConf.CNIVersion)
	}

	cmdDel := func(args *skel.CmdArgs) error {
		var netConf NetConf
		err := json.Unmarshal(args.StdinData, &netConf)
		if err != nil {
			panic(err)
		}
		err = ipam.ExecDel(netConf.IPAM.Type, args.StdinData)
		if err != nil {
			panic(err)
		}
		return nil
	}
	skel.PluginMain(cmdAdd, cmdDel, version.All)
}