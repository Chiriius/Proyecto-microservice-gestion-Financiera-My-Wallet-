

export class AccountDto {

    constructor(
        public id: string ,
        public userId: string,
        public balance: number,
        public accountName: string,
        public createdAt: Date,
        public updatedAt: Date | null = null
    ){}
    static create(object: { [key: string]: any }): [string?, AccountDto?] {
        if (!object.userId) return ["Missing userId"];
        if (!object.accountName) return ["Missing accountName"];
        if (object.balance === undefined || object.balance === null) return ["Missing balance"];

        return [undefined, new AccountDto(object.id,object.userId, object.balance, object.accountName, new Date(), null)];
    }
}