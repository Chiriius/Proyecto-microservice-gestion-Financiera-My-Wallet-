import { Router } from "express";
import { AccountRoutes } from "./controllers/account/routes";




export class AppRoutes {

    static get routes(): Router{
        
        const router = Router();

        router.use('/api', AccountRoutes.routes)

        return router;
    }

}