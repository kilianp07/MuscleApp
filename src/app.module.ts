import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { databaseProviders } from './database/database.provider';
import { UsersModule } from './users/users.module';
import { APP_INTERCEPTOR } from '@nestjs/core';
import { AuthInterceptor } from './interceptor/auth.interceptor';
import { AuthModule } from './auth/auth.module';
import { DatabaseModule } from './database/database.module';
import { UsersService } from './users/users.service';

@Module({
  imports: [AuthModule, UsersModule, DatabaseModule],
  controllers: [AppController],
  providers: [
    AppService,
    {
      provide: APP_INTERCEPTOR,
      useClass: AuthInterceptor,
    },
    ...databaseProviders,
  ],
  exports: [...databaseProviders],
})
export class AppModule {}
