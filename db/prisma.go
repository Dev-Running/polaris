package db

import "github.com/laurentino14/user/prisma"

func UseDB() *prisma.PrismaClient {
	client := prisma.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return nil
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	return client
}

func DeferDB() {

}
