import { Request, Response } from 'express';
import { AccountService, CustomError } from '../../../domain';


export class AccountController  {
    constructor(
        private readonly accountService: AccountService
){}

    private handleError = (error:unknown, res: Response) => {
        if (error instanceof CustomError) {
            res.status(error.statusCode).json({error: error.message});
            return;
        } 
        return res.status(500).json({error: 'Internal Server Error'});

    }

    createAccount = async (req: Request, res: Response) => {
        try {
            const account = await this.accountService.createAccount(req.body);
            res.status(201).json(account);
        } catch (error: any) {
            this.handleError(error, res);
        }
    }
}
