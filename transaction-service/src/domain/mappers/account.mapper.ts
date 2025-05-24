import { Account } from "../../generated/prisma";
import { AcountEntity } from "../entities/account/account.entity";
import {  AccountDto, CreateAccountDto } from "../dto";

export class AccountMapper {
  static toDomain(account: Account): AcountEntity {
    return new AcountEntity(
      account.id,
      account.userId,
      account.balance,
      account.accountName,
      account.createdAt,
      account.updatedAt
    );
  }

  static toPersistence(accountEntity: AcountEntity): { 
    userId: string; 
    balance: number; 
    accountName: string; 
    createdAt: Date; 
    updatedAt: Date | null; 
  } {
    return {
      userId: accountEntity.userId,
      balance: accountEntity.balance,
      accountName: accountEntity.accountName,
      createdAt: accountEntity.createdAt,
      updatedAt: accountEntity.updatedAt,
    };
  }

  static fromCreateDTO(dto: CreateAccountDto): AcountEntity {
    return new AcountEntity(
      '', 
      dto.userId,
      dto.balance,
      dto.accountName,
      new Date(), 
      null
    );
  }

  static toDTO(accountEntity: AcountEntity): AccountDto {
    return {
      id: accountEntity.id,
      userId: accountEntity.userId,
      balance: accountEntity.balance,
      accountName: accountEntity.accountName,
      createdAt: accountEntity.createdAt,
      updatedAt: accountEntity.updatedAt,
    };
  }
}