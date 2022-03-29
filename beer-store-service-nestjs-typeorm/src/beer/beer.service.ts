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

  findAll(search = '', page = 1, pageSize = 10): Promise<Beer[]> {
    return this.beerRepository.query(
      `
      select * from beer 
      where lower(titlebeer) like lower(?) 
      or lower(descriptionbeer) like lower(?)
      limit ? offset ?
    `,
      [`%${search}%`, `%${search}%`, pageSize, (page - 1) * pageSize],
    );
  }

  findOne(id: number): Promise<Beer> {
    return this.beerRepository.findOne(id);
  }
}
