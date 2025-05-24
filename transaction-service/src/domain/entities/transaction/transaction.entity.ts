

export class TransactionEntity {
  constructor(
    public id: string,
    public accountId: string,
    public amount: number,
    public transactionType: [string],
    public category: string,
    public description: string,
    public createdAt: Date,
    public updatedAt: Date
  ) {}
}