/*
  Warnings:

  - You are about to drop the column `cellphone` on the `users` table. All the data in the column will be lost.

*/
-- DropIndex
DROP INDEX "users_cellphone_key";

-- AlterTable
ALTER TABLE "users" DROP COLUMN "cellphone";
