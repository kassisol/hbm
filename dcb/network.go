package dcb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/docker/engine-api/types"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/pkg/cmdbuilder"
	"github.com/kassisol/hbm/pkg/utils"
)

func NetworkList(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("network")
	cmd.Add("ls")

	cmd.GetParams(req.RequestURI)

	if len(cmd.Params) > 0 {
		if _, ok := cmd.Params["filters"]; ok {
			var v map[string]map[string]bool

			err := json.Unmarshal([]byte(cmd.Params["filters"][0]), &v)
			if err != nil {
				panic(err)
			}

			var r []string

			for k, val := range v {
				r = append(r, k)

				for ka, _ := range val {
					r = append(r, ka)
				}
			}

			cmd.Add(fmt.Sprintf("--filter \"%s=%s\"", r[0], r[1]))
		}
	}

	return cmd.String()
}

func NetworkInspect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("network")
	cmd.Add("inspect")

	_, urlPath := utils.GetURIInfo(req)

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}

func NetworkCreate(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("network")
	cmd.Add("create")

	nc := &types.NetworkCreateRequest{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(nc); err != nil {
			panic(err)
		}
	}

	if len(nc.Driver) > 0 {
		cmd.Add(fmt.Sprintf("--driver %s", nc.Driver))
	}

	if nc.Internal {
		cmd.Add("--internal")
	}

	if nc.EnableIPv6 {
		cmd.Add("--ipv6")
	}

	// IPAM
	if len(nc.IPAM.Driver) > 0 {
		cmd.Add(fmt.Sprintf("--ipam-driver=%s", nc.IPAM.Driver))
	}

	if len(nc.IPAM.Options) > 0 {
		for k, v := range nc.IPAM.Options {
			cmd.Add(fmt.Sprintf("--ipam-opt %s=%s", k, v))
		}
	}

	if len(nc.IPAM.Config) > 0 {
		for _, c := range nc.IPAM.Config {
			if len(c.Subnet) > 0 {
				cmd.Add(fmt.Sprintf("--subnet %s", c.Subnet))
			}

			if len(c.IPRange) > 0 {
				cmd.Add(fmt.Sprintf("--ip-range %s", c.IPRange))
			}

			if len(c.Gateway) > 0 {
				cmd.Add(fmt.Sprintf("--gateway %s", c.Gateway))
			}

			if len(c.AuxAddress) > 0 {
				for k, v := range c.AuxAddress {
					cmd.Add(fmt.Sprintf("--aux-address %s=%s", k, v))
				}
			}
		}
	}

	// Options
	if len(nc.Options) > 0 {
		for k, v := range nc.Options {
			cmd.Add(fmt.Sprintf("--opt %s=%s", k, v))
		}
	}

	if len(nc.Labels) > 0 {
		for k, v := range nc.Labels {
			cmd.Add(fmt.Sprintf("--label %s=%s", k, v))
		}
	}

	if len(nc.Name) > 0 {
		cmd.Add(nc.Name)
	}

	return cmd.String()
}

func NetworkConnect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("network")
	cmd.Add("connect")

	nc := &types.NetworkConnect{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(nc); err != nil {
			panic(err)
		}
	}

	if len(nc.EndpointConfig.IPAMConfig.IPv4Address) > 0 {
		cmd.Add(fmt.Sprintf("--ip %s", nc.EndpointConfig.IPAMConfig.IPv4Address))
	}

	if len(nc.EndpointConfig.IPAMConfig.IPv6Address) > 0 {
		cmd.Add(fmt.Sprintf("--ip6 %s", nc.EndpointConfig.IPAMConfig.IPv6Address))
	}

	if len(nc.EndpointConfig.Links) > 0 {
		for _, v := range nc.EndpointConfig.Links {
			cmd.Add(fmt.Sprintf("--link %s", v))
		}
	}

	if len(nc.EndpointConfig.Aliases) > 0 {
		for _, v := range nc.EndpointConfig.Aliases {
			cmd.Add(fmt.Sprintf("--alias %s", v))
		}
	}

	// Network ID
	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	// Container
	if len(nc.Container) > 0 {
		cmd.Add(nc.Container)
	}

	return cmd.String()
}

func NetworkDisconnect(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("network")
	cmd.Add("disconnect")

	nd := &types.NetworkDisconnect{}

	if req.RequestBody != nil {
		if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(nd); err != nil {
			panic(err)
		}
	}

	if nd.Force {
		cmd.Add("--force")
	}

	// Network ID
	_, urlPath := utils.GetURIInfo(req)
	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	// Container
	if len(nd.Container) > 0 {
		cmd.Add(nd.Container)
	}

	return cmd.String()
}

func NetworkRemove(req authorization.Request, re *regexp.Regexp) string {
	cmd := cmdbuilder.New("network")
	cmd.Add("rm")

	_, urlPath := utils.GetURIInfo(req)

	cmd.Add(re.FindStringSubmatch(urlPath)[1])

	return cmd.String()
}
