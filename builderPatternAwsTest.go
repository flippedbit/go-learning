package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Filter struct {
	f []*ec2.Filter
}

func (f *Filter) ByRunning() *Filter {
	f.f = append(f.f, &ec2.Filter{
		Name: aws.String("instance-state-name"),
		Values: []*string{
			aws.String("running"),
		},
	})
	return f
}

func (f *Filter) BySubnet(s string) *Filter {
	ss := strings.Split(s, ",")
	f.f = append(f.f, &ec2.Filter{
		Name:   aws.String("subnet-id"),
		Values: aws.StringSlice(ss),
	})
	return f
}

func (f *Filter) ByTags(s string) *Filter {

	return f
}

func (f *Filter) ByAvailabilityZone(s string) *Filter {
	ss := strings.Split(s, ",")
	f.f = append(f.f, &ec2.Filter{
		Name:   aws.String("availability-zone"),
		Values: aws.StringSlice(ss),
	})
	return f
}

func (f *Filter) ByInstance(s string) *Filter {
	ss := strings.Split(s, ",")
	f.f = append(f.f, &ec2.Filter{
		Name:   aws.String("instance-id"),
		Values: aws.StringSlice(ss),
	})
	return f
}

func (f *Filter) Build() []*ec2.Filter {
	return f.f
}

func main() {
	// new way to construct aws Filter for gathering instances, just as an example using builder pattern
	filter := &Filter{}
	f := filter.ByRunning().BySubnet("abc,bcd").ByInstance("i-12345,i-56789").ByAvailabilityZone("us-east-1a").Build()
	fmt.Println(f)
	// output:
	// [{
	//   Name: "instance-state-name",
	//   Values: ["running"]
	// } {
	//   Name: "subnet-id",
	//   Values: ["abc","bcd"]
	// } {
	// 	 Name: "availability-zone",
	// 	 Values: ["us-east-1a"]
	// }]

	// previous Filter used in app that works for comparison
	e := []*ec2.Filter{
		{
			Name: aws.String("subnet-id"),
			Values: []*string{
				aws.String("subnet-abc"),
			},
		},
		{
			Name: aws.String("instance-state-name"),
			Values: []*string{
				aws.String("running"),
			},
		},
	}
	fmt.Println(e)
	// output:
	// [{
	// 	Name: "subnet-id",
	// 	Values: ["subnet-abc"]
	// } {
	// 	Name: "instance-state-name",
	// 	Values: ["running"]
	// }]
}
