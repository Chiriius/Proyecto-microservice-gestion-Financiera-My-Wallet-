import { AccountGrpcController } from "./account.grpc.controller";
import { AccountRepository } from "../../../domain/repositories/account.repository";
import { AccountService } from "../../../domain/services/account.service";

export class AccountGrpcRoutes {
    static get grpcService() {
        const accountRepository = new AccountRepository();
        const accountService = new AccountService(accountRepository);
        const accountGrpcController = new AccountGrpcController(accountService);

        return {
            CreateAccount: accountGrpcController.createAccount,
        };
    }
}