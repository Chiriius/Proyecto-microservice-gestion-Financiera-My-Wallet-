
import { AccountGrpcRoutes } from "./controllers/account/account.routes.grpc";

export class AppGrpcRoutes {
    static get grpcServices() {
        return {
            ...AccountGrpcRoutes.grpcService,

        };
    }
}