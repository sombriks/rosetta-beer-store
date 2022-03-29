import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { BeerController } from './beer.controller';
import { Beer } from './beer.entity';
import { BeerService } from './beer.service';

@Module({
  imports: [TypeOrmModule.forFeature([Beer])],
  exports: [TypeOrmModule],
  providers: [BeerService],
  controllers: [BeerController],
})
export class BeerModule {}
