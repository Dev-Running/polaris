/*
  Warnings:

  - You are about to drop the column `course_id` on the `enrollments` table. All the data in the column will be lost.
  - You are about to drop the column `user_id` on the `enrollments` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "enrollments" DROP COLUMN "course_id",
DROP COLUMN "user_id";
