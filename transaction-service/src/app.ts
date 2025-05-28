import { envs, loadSecrets } from "./config/envs";
import { prisma } from "./domain/entities/prisma.client";
import { AppRoutes } from "./presentation/routes";
import { AppGrpcRoutes } from "./presentation/routes.grpc";
import { Server } from "./presentation/server";

(async () => {
    await main();
})();

async function main() {
    try {
        await loadSecrets();
        await prisma.$connect();
        console.log("Conexión a PostgreSQL establecida correctamente.");

        new Server({
            port: envs.PORT, 
            grpcPort: envs.GRPC_PORT, 
            routes: AppRoutes.routes,
            protoPath: __dirname + "/presentation/proto/account.proto",
            grpcService: AppGrpcRoutes.grpcServices,
        }).start();
    } catch (error) {
        console.error("Error al iniciar la aplicación:", error);
        process.exit(1); 
    }
}