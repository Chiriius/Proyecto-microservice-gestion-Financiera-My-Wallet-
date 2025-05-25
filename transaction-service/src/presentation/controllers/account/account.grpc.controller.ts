import { AccountService } from "../../../domain";
import * as grpc from '@grpc/grpc-js';



export class AccountGrpcController {
    constructor(
        private readonly accountService : AccountService
    ){}

    createAccount = async (call: any, callback: any) => {
        try {
            const { userId, balance, accountName } = call.request;
            const account = await this.accountService.createAccount({ userId, balance, accountName });
            callback(null, account);
        } catch (error) {
            console.error("Error en CreateAccount:", error);
            callback({
                code: grpc.status.INTERNAL,
                message: "Error al crear la cuenta",
            });
        }
    }
}