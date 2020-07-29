// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pb

import (
	client "github.com/NathanBaulch/protoc-gen-cobra/client"
	iocodec "github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	proto "github.com/golang/protobuf/proto"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
)

func CyclicalClientCommand(cfgs ...*client.Config) *cobra.Command {
	cfg := client.DefaultConfig
	if len(cfgs) > 0 {
		cfg = cfgs[0]
	}
	cmd := &cobra.Command{
		Use:   "cyclical",
		Short: "Cyclical service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	d := &client.Dialer{Config: cfg}
	cmd.AddCommand(
		_CyclicalTestCommand(d),
	)
	return cmd
}

func _CyclicalTestCommand(d *client.Dialer) *cobra.Command {
	req := &Foo{
		Bar1: &Bar{},
		Bar2: &Bar{},
	}

	cmd := &cobra.Command{
		Use:   "test",
		Short: "Test RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return d.RoundTrip(cmd.Context(), func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewCyclicalClient(cc)
				v := &Foo{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Test(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Bar1.Value, "bar1-value", "", "")
	cmd.PersistentFlags().StringVar(&req.Bar2.Value, "bar2-value", "", "")
	cmd.PersistentFlags().StringVar(&req.Value, "value", "", "")

	return cmd
}
