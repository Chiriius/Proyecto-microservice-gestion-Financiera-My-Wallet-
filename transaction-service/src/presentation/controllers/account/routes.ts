import { Router } from "express";
import { AccountController } from "./account.controller";
import { AccountRepository, AccountService } from "../../../domain";


export class AccountRoutes {
    static get routes():Router{
        const router = Router();
        const accountRepository = new AccountRepository();
        const accountService =  new AccountService(accountRepository);
        const accountController = new AccountController(accountService);

        router.post('/accounts', accountController.createAccount);

        return router;
    }
}