// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pb

import (
	client "github.com/NathanBaulch/protoc-gen-cobra/client"
	iocodec "github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	proto "github.com/golang/protobuf/proto"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
)

func BankClientCommand(cfgs ...*client.Config) *cobra.Command {
	cfg := client.DefaultConfig
	if len(cfgs) > 0 {
		cfg = cfgs[0]
	}
	cmd := &cobra.Command{
		Use:   "bank",
		Short: "Bank service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	d := &client.Dialer{Config: cfg}
	cmd.AddCommand(
		_BankDepositCommand(d),
	)
	return cmd
}

func _BankDepositCommand(d *client.Dialer) *cobra.Command {
	req := &DepositRequest{
		ClusterWithNamespaces: &DepositRequest_ClusterWithNamespaces{
			Cluster: &Cluster{},
		},
	}

	cmd := &cobra.Command{
		Use:   "deposit",
		Short: "Deposit RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return d.RoundTrip(cmd.Context(), func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewBankClient(cc)
				v := &DepositRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Deposit(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Environment, "environment", "", "")
	cmd.PersistentFlags().StringVar(&req.Parent, "parent", "", "")
	cmd.PersistentFlags().StringVar(&req.Tenant, "tenant", "", "")

	return cmd
}
