import vault from "node-vault";

const VAULT_ADDR = process.env.VAULT_ADDRESS || "http://127.0.0.1:8200";
const VAULT_TOKEN = process.env.VAULT_TOKEN || "dev-only-token";

const vaultClient = vault({
  endpoint: VAULT_ADDR,
  token: VAULT_TOKEN,
});

export const getSecrets = async (path: string) => {
  try {
    const response = await vaultClient.read(path);
    return response.data.data;
  } catch (error) {
    console.error("Error al obtener secretos desde Vault:", error);
    throw error;
  }
};