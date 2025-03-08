/*
  Warnings:

  - You are about to drop the column `review_id` on the `reviews` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE `reviews` DROP FOREIGN KEY `reviews_review_id_fkey`;

-- AlterTable
ALTER TABLE `reviews` DROP COLUMN `review_id`;

-- AddForeignKey
ALTER TABLE `reviews` ADD CONSTRAINT `reviews_id_fkey` FOREIGN KEY (`id`) REFERENCES `readings`(`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
