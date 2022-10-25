/*
  Warnings:

  - You are about to drop the column `roomsId` on the `Messages` table. All the data in the column will be lost.
  - You are about to drop the column `sent_at` on the `Messages` table. All the data in the column will be lost.
  - You are about to drop the column `socket` on the `users` table. All the data in the column will be lost.
  - You are about to drop the `contacts` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `rooms` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "Messages" DROP CONSTRAINT "Messages_roomsId_fkey";

-- DropForeignKey
ALTER TABLE "contacts" DROP CONSTRAINT "contacts_userId_fkey";

-- DropForeignKey
ALTER TABLE "rooms" DROP CONSTRAINT "rooms_userId_fkey";

-- AlterTable
ALTER TABLE "Messages" DROP COLUMN "roomsId",
DROP COLUMN "sent_at",
ADD COLUMN     "created_at" DATE NOT NULL DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN     "userId" TEXT;

-- AlterTable
ALTER TABLE "users" DROP COLUMN "socket",
ADD COLUMN     "socketID" TEXT;

-- DropTable
DROP TABLE "contacts";

-- DropTable
DROP TABLE "rooms";

-- AddForeignKey
ALTER TABLE "Messages" ADD CONSTRAINT "Messages_userId_fkey" FOREIGN KEY ("userId") REFERENCES "users"("id") ON DELETE SET NULL ON UPDATE CASCADE;
