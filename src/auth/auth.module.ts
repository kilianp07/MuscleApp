import { Module } from '@nestjs/common';
import { AuthService } from './auth.service';
import { UsersModule } from '../users/users.module';
import { LocalStrategy } from './strategy/local.strategy';
import { PassportModule } from '@nestjs/passport';
import { JwtModule } from '@nestjs/jwt';
import { JwtStrategy } from './strategy/jwt.strategy';
import { jwtConstants } from './constants';
import { AuthController } from './auth.controller';
import { UsersService } from 'src/users/users.service';

@Module({
  providers: [AuthService, LocalStrategy, JwtStrategy, UsersService],
  imports: [
    UsersModule,
    PassportModule,
    JwtModule.register({
      secret: jwtConstants.secret,
      signOptions: { expiresIn: '24h' },
    }),
  ],
  controllers: [AuthController],
  exports: [AuthService],
})
export class AuthModule {}
