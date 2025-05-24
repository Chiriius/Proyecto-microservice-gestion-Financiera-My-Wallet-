

export class AcountEntity {
  constructor(
    public id: string,
    public userId: string,
    public balance: number,
    public accountName: string,
    public createdAt: Date ,
    public updatedAt: Date| null
  ) {}
}