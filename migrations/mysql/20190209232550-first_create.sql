
-- +migrate Up

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema emukone
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Table `users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128) NOT NULL,
  `identifier` VARCHAR(45) NOT NULL,
  `icon_url` VARCHAR(1024) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted` TINYINT NOT NULL DEFAULT 0,
  `header_image_url` VARCHAR(1024) NOT NULL,
  `meta_description` VARCHAR(512) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `user_identifier` (`identifier` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `artist_users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `artist_users` (
  `user_id` BIGINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `identifier` VARCHAR(45) NOT NULL DEFAULT 'artist',
  `why_description` TEXT NOT NULL,
  `how_description` TEXT NOT NULL,
  `free_description` TEXT NOT NULL,
  PRIMARY KEY (`user_id`),
  INDEX `artist_identifier` (`identifier` ASC),
  CONSTRAINT `artist_user_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `events`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `events` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128) NOT NULL,
  `start_time` DATETIME NOT NULL,
  `end_time` DATETIME NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `canceled` TINYINT NOT NULL DEFAULT 0,
  `why_description` TEXT NOT NULL,
  `free_description` TEXT NOT NULL,
  `place` TEXT NOT NULL,
  `ticket_price` TEXT NOT NULL,
  `header_image_url` VARCHAR(1024) NOT NULL,
  `meta_description` VARCHAR(512) NOT NULL,
  `open_time` DATETIME NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `invites`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `invites` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `artist_user_id` BIGINT NOT NULL,
  `event_id` BIGINT NOT NULL,
  `description` TEXT NOT NULL,
  `status` INT NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `canceled` TINYINT NOT NULL DEFAULT 0,
  `inviter_user_id` BIGINT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `invite_artist_to_artist_id_idx` (`artist_user_id` ASC),
  INDEX `invite_artist_event_id_idx` (`event_id` ASC),
  INDEX `invite_inviter_user_id_idx` (`inviter_user_id` ASC),
  CONSTRAINT `invite_artist_artist_id`
    FOREIGN KEY (`artist_user_id`)
    REFERENCES `artist_users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `invite_artist_event_id`
    FOREIGN KEY (`event_id`)
    REFERENCES `events` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `invite_inviter_user_id`
    FOREIGN KEY (`inviter_user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `events_related_users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `events_related_users` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT NOT NULL,
  `event_id` BIGINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `erl_event_id_idx` (`event_id` ASC),
  INDEX `erl_user_id_idx` (`user_id` ASC),
  CONSTRAINT `erl_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `erl_event_id`
    FOREIGN KEY (`event_id`)
    REFERENCES `events` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `booker_users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `booker_users` (
  `user_id` BIGINT NOT NULL,
  `identifier` VARCHAR(45) NOT NULL DEFAULT 'booker',
  `why_description` TEXT NOT NULL,
  `how_description` TEXT NOT NULL,
  `free_description` TEXT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  INDEX `booker_identifier` (`identifier` ASC),
  CONSTRAINT `user_id_booker_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `user_tags`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `user_tags` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `user_tags_related_users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `user_tags_related_users` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT NOT NULL,
  `user_tag_id` BIGINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `tru_user_id_idx` (`user_id` ASC),
  INDEX `tru_tag_id_idx` (`user_tag_id` ASC),
  CONSTRAINT `tru_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `tru_tag_id`
    FOREIGN KEY (`user_tag_id`)
    REFERENCES `user_tags` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `artist_members`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `artist_members` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128) NOT NULL,
  `part` VARCHAR(128) NOT NULL,
  `artist_user_id` BIGINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `artist_member_user_id_idx` (`artist_user_id` ASC),
  CONSTRAINT `artist_member_user_id`
    FOREIGN KEY (`artist_user_id`)
    REFERENCES `artist_users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `locations`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `locations` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `created_at` DATETIME NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `locations_related_users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `locations_related_users` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT NOT NULL,
  `location_id` BIGINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `lru_user_id_idx` (`user_id` ASC),
  INDEX `lru_location_id_idx` (`location_id` ASC),
  CONSTRAINT `lru_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `lru_location_id`
    FOREIGN KEY (`location_id`)
    REFERENCES `locations` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `website_urls`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `website_urls` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `type` VARCHAR(45) NOT NULL,
  `url` VARCHAR(1024) NOT NULL,
  `user_id` BIGINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `type_and_user_unique` (`type` ASC, `user_id` ASC),
  INDEX `website_urls_user_id_idx` (`user_id` ASC),
  CONSTRAINT `website_urls_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ichioshis`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ichioshis` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `service` VARCHAR(45) NOT NULL,
  `url` VARCHAR(1024) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `artist_user_id` BIGINT NOT NULL,
  `embed` VARCHAR(128) NOT NULL,
  `title` VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `ichioshis_artist_user_id_idx` (`artist_user_id` ASC),
  CONSTRAINT `ichioshis_artist_user_id`
    FOREIGN KEY (`artist_user_id`)
    REFERENCES `artist_users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `external_service_uids`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `external_service_uids` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `service` VARCHAR(45) NOT NULL,
  `uid` VARCHAR(128) NOT NULL,
  `user_id` BIGINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `esu_user_id_idx` (`user_id` ASC),
  UNIQUE INDEX `unique_esu_user_id_service` (`service` ASC, `user_id` ASC),
  CONSTRAINT `esu_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `user_favorites`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `user_favorites` (
  `user_id` BIGINT NOT NULL AUTO_INCREMENT,
  `counts` BIGINT NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `user_favorite_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `event_favorites`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `event_favorites` (
  `user_id` BIGINT NOT NULL,
  `counts` BIGINT NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `favorite_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `events` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `user_follow`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `user_follow` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `follow_user_id` BIGINT NOT NULL,
  `user_id` BIGINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `user_follow_user_id_idx` (`follow_user_id` ASC),
  INDEX `user_follow_user_id_idx1` (`user_id` ASC),
  CONSTRAINT `user_followed_user_id`
    FOREIGN KEY (`follow_user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `user_follow_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `event_follows`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `event_follows` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `event_id` BIGINT NOT NULL,
  `user_id` BIGINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `event_follow_event_id_idx` (`event_id` ASC),
  INDEX `event_follow_user_id_idx` (`user_id` ASC),
  CONSTRAINT `event_follow_event_id`
    FOREIGN KEY (`event_id`)
    REFERENCES `events` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `event_follow_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `user_images`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `user_images` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `image_url` VARCHAR(1024) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_id` BIGINT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `user_images_user_id_idx` (`user_id` ASC),
  CONSTRAINT `user_images_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `event_images`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `event_images` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `image_url` VARCHAR(1024) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `event_id` BIGINT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `event_image_event_id_idx` (`event_id` ASC),
  CONSTRAINT `event_image_event_id`
    FOREIGN KEY (`event_id`)
    REFERENCES `events` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

-- +migrate Down
