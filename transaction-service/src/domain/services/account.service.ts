import { AccountDto, CreateAccountDto } from "../dto";
import { AccountRepository } from "../repositories/account.repository";


export class AccountService {

  constructor(
    private readonly accountRepository: AccountRepository
    ){}

  async createAccount(dto: CreateAccountDto): Promise<AccountDto> {
    if (dto.balance < 0) {
      throw new Error('El balance no puede ser negativo');
    }

    return await this.accountRepository.createAccount(dto);
  }

}