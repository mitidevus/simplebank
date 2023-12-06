DROP TABLE IF EXIST "verify_emails" CASCADE;

ALTER TABLE "users" DROP COLUMN "is_email_verified";