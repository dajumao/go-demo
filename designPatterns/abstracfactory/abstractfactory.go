package abstracfactory

import "fmt"

type OrderMainDAO interface {
	SaveOrderMain()
}

type OrderDetailDAO interface {
	SaveOrderDetail()
}

type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

type RDBMainDAO struct {

}

func (*RDBMainDAO)SaveOrderMain()  {
	fmt.Print("rdb main save\n")
}

type RDBDetailDAO struct {

}

func (*RDBDetailDAO)SaveOrderDetail()  {
	fmt.Print("rdb detail save\n")
}

type RDBDAOFactory struct {

}

func (*RDBDAOFactory)CreateOrderMainDAO() OrderMainDAO  {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory)CreateOrderDetailDAO() OrderDetailDAO  {
	return &RDBDetailDAO{}
}

type XMLMainDAO struct {

}

func (*XMLMainDAO)SaveOrderMain()  {
	fmt.Print("xml main save\n")
}

type XMLDetailDAO struct {

}

func (*XMLDetailDAO)SaveOrderDetail()  {
	fmt.Print("xml detail save")
}

type XMLDAOFactory struct {

}

func (*XMLDAOFactory)CreateOrderMainDAO() OrderMainDAO  {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory)CreateOrderDetailDAO() OrderDetailDAO  {
	return &XMLDetailDAO{}
}