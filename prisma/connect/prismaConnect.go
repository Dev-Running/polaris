package connect

import (
	"github.com/laurentino14/user/prisma"
)

type DB struct {
	Client *prisma.PrismaClient
}

func NewPrismaConnect() *DB {

	prismaConnect := &DB{}
	prismaConnect.Client = prisma.NewClient()

	if err := prismaConnect.Client.Prisma.Connect(); err != nil {
		panic("Banco de dados")
	}

	return prismaConnect
}

func (p *DB) Disconnect() {

	if err := p.Client.Prisma.Disconnect(); err != nil {
		panic(err)
	}

}
