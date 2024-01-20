use cookbook;

CREATE TABLE `recipe` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT NOT NULL,
    `direction` TEXT NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=INNODB AUTO_INCREMENT=1540 DEFAULT CHARSET=utf8;

CREATE TABLE `ingredient` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=INNODB AUTO_INCREMENT=1540 DEFAULT CHARSET=utf8;

CREATE TABLE `unit` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=INNODB AUTO_INCREMENT=1540 DEFAULT CHARSET=utf8;

CREATE TABLE `recipe_ingredient` (
    `recipe_id` BIGINT UNSIGNED NOT NULL,
    `ingredient_id` BIGINT UNSIGNED NOT NULL,
    `unit_id` BIGINT UNSIGNED NOT NULL,
    `quantity` DECIMAL(15,4) NOT NULL,
    `note` TEXT NOT NULL
) ENGINE=INNODB AUTO_INCREMENT=1540 DEFAULT CHARSET=utf8;

ALTER TABLE `recipe_ingredient` 
ADD CONSTRAINT fk_recipe_ingredient_recipe
FOREIGN KEY (`recipe_id`) 
REFERENCES `recipe`(`id`);

ALTER TABLE `recipe_ingredient` 
ADD CONSTRAINT fk_recipe_ingredient_ingredient
FOREIGN KEY (`ingredient_id`) 
REFERENCES `ingredient`(`id`);

ALTER TABLE `recipe_ingredient` 
ADD CONSTRAINT fk_recipe_ingredient_unit
FOREIGN KEY (`unit_id`) 
REFERENCES `unit`(`id`);