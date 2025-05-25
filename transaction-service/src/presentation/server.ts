import express, { Router } from 'express';
import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';

interface Options {
    port?: number;
    routes: Router;
    grpcPort?: number;
    protoPath: string;
    grpcService: any;
}

export class Server {
    public readonly app = express();
    private readonly port: number;
    private readonly grpcPort: number;
    private readonly routes: Router;
    private readonly protoPath: string;
    private readonly grpcService: any;

    constructor(option: Options) {
        const { port = 3100, grpcPort = 50051, routes, protoPath, grpcService } = option;
        this.port = port;
        this.grpcPort = grpcPort;
        this.routes = routes;
        this.protoPath = protoPath;
        this.grpcService = grpcService;
    }

    async start() {
        // Iniciar servidor HTTP
        this.app.use(express.json());
        this.app.use(express.urlencoded({ extended: true }));
        this.app.use(this.routes);

        this.app.listen(this.port, () => {
            console.log(`Servidor HTTP ejecutándose en el puerto ${this.port}`);
        });

        // Iniciar servidor gRPC
        const packageDefinition = protoLoader.loadSync(this.protoPath, {
            keepCase: true,
            longs: String,
            enums: String,
            defaults: true,
            oneofs: true,
        });
        
        const grpcPackage = grpc.loadPackageDefinition(packageDefinition) as any;
               
        if (!grpcPackage.account || !grpcPackage.account.AccountService) {
            throw new Error("El servicio gRPC 'AccountService' no está definido. Verifica el archivo .proto y su carga.");
        }
        
        const server = new grpc.Server();
        
        server.addService(grpcPackage.account.AccountService.service, this.grpcService);
        
        server.bindAsync(`0.0.0.0:${this.grpcPort}`, grpc.ServerCredentials.createInsecure(), () => {
            console.log(`Servidor gRPC ejecutándose en el puerto ${this.grpcPort}`);
        });
    }
}