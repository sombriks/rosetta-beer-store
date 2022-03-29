import { Controller, Get, Param, Query, Req } from '@nestjs/common';
import { Beer } from './beer.entity';
import { BeerService } from './beer.service';

@Controller('beer')
export class BeerController {
  constructor(private readonly beerService: BeerService) {}

  @Get('list')
  async list(
    @Query('search') search: string,
    @Query('page') page: number,
    @Query('pageSize') pageSize: number,
  ): Promise<Beer[]> {
    return this.beerService.findAll(search,page,pageSize);
  }

  @Get(':idbeer')
  async find(@Param('idbeer') idbeer: number): Promise<Beer> {
    return this.beerService.findOne(idbeer);
  }
}
