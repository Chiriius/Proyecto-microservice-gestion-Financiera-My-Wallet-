import { getSecrets } from "./vault.config";

let secrets: Record<string, any> = {};


export const loadSecrets = async () => {
  try {
    secrets = await getSecrets("secret/data/mywallet");

    process.env.DATABASE_URL_POSTGRESQL_TRANSACTIONS = secrets.DATABASE_URL_POSTGRESQL_TRANSACTIONS;
    process.env.GRPC_PORT = secrets.GRPC_PORT;
    
  } catch (error) {
    console.error("Error al cargar secretos:", error);
    process.exit(1); 
  }
};

export const envs = {
  get PORT() {
    return parseInt(secrets.PORTSERVICE_TRANSACTION, 10);
  },
  get DATABASE_URL_POSTGRESQL_TRANSACTIONS() {
    return secrets.DATABASE_URL_POSTGRESQL_TRANSACTIONS;
  },
  get GRPC_PORT() {
    return parseInt(secrets.GRPC_PORT, 10);
  },
};