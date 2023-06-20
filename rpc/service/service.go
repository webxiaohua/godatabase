package service

import "errors"

type Args struct {
	A,B int
}

type Quotient struct {
	Quo,Rem int
}

type Service int

func (*Service) Multiply(args *Args,reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (*Service) Divide(args *Args,quo *Quotient) error {
	if args.B == 0 {
		return errors.New("devide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}