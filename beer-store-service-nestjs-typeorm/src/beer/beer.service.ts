import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { Beer } from './beer.entity';

@Injectable()
export class BeerService {
  constructor(
    @InjectRepository(Beer)
    private beerRepository: Repository<Beer>,
  ) {}

  findAll(): Promise<Beer[]> {
    return this.beerRepository.find();
  }

  findOne(id: string): Promise<Beer> {
    return this.beerRepository.findOne(id);
  }
}
