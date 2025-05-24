import { envs } from "./config/envs";
import { prisma } from "./domain/entities/prisma.client";
import { AppRoutes } from "./presentation/routes";
import { Server } from "./presentation/server";

(async () => {
    await main();
})();

async function main() {
    try {
        await prisma.$connect();
        console.log('Conexión a PostgreSQL establecida correctamente.');

        new Server({
            port: envs.PORT,
            routes: AppRoutes.routes,
        }).start();
    } catch (error) {
        console.error('Error al iniciar la aplicación:', error);
        process.exit(1); 
    }
}