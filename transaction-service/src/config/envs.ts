import "dotenv/config";
import * as dotenv from 'dotenv';
import { get } from "env-var";

dotenv.config({path: '../.env'});



export const envs = {
    PORT: get('PORTSERVICE_TRANSACTION').required().asPortNumber(),
    DATABASE_URL_POSTGRESQL_TRANSACTIONS: get('DATABASE_URL_POSTGRESQL_TRANSACTIONS').required().asString(),
}