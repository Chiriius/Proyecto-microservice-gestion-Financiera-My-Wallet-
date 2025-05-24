
export class CreateAccountDto{

    constructor(
        public userId: string,
        public balance: number,
        public accountName: string,
    ){}

    static create(object:{[key:string]:any}): [string?,CreateAccountDto?] {
        if (!object.userId)      return ["Missing userId"];
        if (!object.accountName) return ["Missing accountName"];
        if (!object.balance && object.balance !== 0) return ["Missing balance"];
    
        return [undefined, new CreateAccountDto(object.userId, object.balance, object.accountName)];
    }
}
