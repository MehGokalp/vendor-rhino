CREATE TABLE `currency` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(4) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `currency_code_UNIQUE` (`code` ASC));

CREATE TABLE `card` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `balance` INT NOT NULL,
  `currency_id` INT NOT NULL,
  `activation_date` DATE NOT NULL,
  `expire_date` DATE NOT NULL,
  `reference` VARCHAR(255) NOT NULL,
  `card_number` VARCHAR(255) NOT NULL,
  `cvc` VARCHAR(255) NOT NULL,
  `active` TINYINT(1) NULL DEFAULT 1,
  `deleted_at` DATE NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `card_card_number_UNIQUE` (`card_number` ASC),
  INDEX `card_currency_id_idx` (`currency_id` ASC),
  CONSTRAINT `currency_idx`
      FOREIGN KEY (`currency_id`)
          REFERENCES `currency` (`id`)
          ON DELETE CASCADE
          ON UPDATE CASCADE);
