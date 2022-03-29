import { Controller, Get, Req } from '@nestjs/common';
import { Beer } from './beer.entity';
import { BeerService } from './beer.service';
import { Request } from 'express';

@Controller('beer')
export class BeerController {
  constructor(private readonly beerService: BeerService) {}

  @Get('list')
  async list(@Req() request: Request): Promise<Beer[]> {
    return this.beerService.findAll();
  }
}
