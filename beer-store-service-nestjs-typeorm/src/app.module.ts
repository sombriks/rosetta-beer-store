import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';

import { TypeOrmModule } from '@nestjs/typeorm';
import { BeerModule } from './beer/beer.module';

@Module({
  imports: [TypeOrmModule.forRoot(), BeerModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
