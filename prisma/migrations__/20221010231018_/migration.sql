/*
  Warnings:

  - Made the column `updated_at` on table `enrollments` required. This step will fail if there are existing NULL values in that column.
  - Added the required column `description` to the `lessons` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "enrollments" ALTER COLUMN "updated_at" SET NOT NULL;

-- AlterTable
ALTER TABLE "lessons" ADD COLUMN     "description" TEXT NOT NULL;
