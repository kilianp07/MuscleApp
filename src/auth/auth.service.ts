import { Injectable } from '@nestjs/common';
import { UsersService } from '../users/users.service';
import { JwtService } from '@nestjs/jwt';
import * as bcrypt from 'bcrypt';

@Injectable()
export class AuthService {
  constructor(
    private usersService: UsersService,
    private jwtService: JwtService,
  ) {}

  async validateUser(email: string, pass: string): Promise<any> {
    const user = await this.usersService.findByEmail(email);

    bcrypt.compare(pass, user.password, function (err, result) {
      if (err) {
        throw err;
      }
      return result;
    });

    return null;
  }

  async login(user: any) {
    // console.log(user.user);
    const payload = {
      user: {
        id: user.user.id,
        email: user.user.email,
        name: user.user.name,
        created_at: user.user.created_at,
        updated_at: user.user.updated_at,
      },
    };
    // console.log({payload});
    return {
      access_token: this.jwtService.sign(payload),
    };
  }

  async register(data) {
    data.password = await bcrypt.hash(data.password, 10);
    const response = await this.usersService.create(data);
    if (response) {
      const { password, ...result } = response;
      return result;
    }
  }

  decodeToken(token): any {
    return this.jwtService.decode(token);
  }
}
