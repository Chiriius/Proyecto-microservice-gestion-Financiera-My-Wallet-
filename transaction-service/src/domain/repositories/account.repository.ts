import { AccountDto, CreateAccountDto } from "../dto";
import { prisma } from "../entities";
import { AccountMapper } from "../mappers/account.mapper";



export class AccountRepository {
    
    async createAccount(dto: CreateAccountDto): Promise<AccountDto> {
        try {
          const accountEntity = AccountMapper.fromCreateDTO(dto);
          const account = await prisma.account.create({
            data: AccountMapper.toPersistence(accountEntity),
          });
          return AccountMapper.toDTO(AccountMapper.toDomain(account));
        } catch (error) {
          console.error("Error al crear la cuenta:", error);
          throw new Error("No se pudo crear la cuenta. Por favor, inténtelo de nuevo.");
        }
      }
  
    async getAccountById(id: string): Promise<AccountDto | null> {
      const account = await prisma.account.findUnique({
        where: { id },
      });
      return account ? AccountMapper.toDTO(AccountMapper.toDomain(account)) : null;
    }
  }