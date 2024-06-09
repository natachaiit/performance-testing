import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { Item } from './item.entity';

@Injectable()
export class ItemsService {
  constructor(
    @InjectRepository(Item)
    private itemsRepository: Repository<Item>,
  ) {}

  async createItem(name: string): Promise<Item> {
    const newItem = this.itemsRepository.create({ name });
    return this.itemsRepository.save(newItem);
  }

  async getItems(): Promise<Item[]> {
    return this.itemsRepository.find();
  }

  async getItemById(id: number): Promise<Item> {
    const item = await this.itemsRepository.findOneBy({ id });
    if (!item) {
      throw new NotFoundException('Item not found');
    }
    return item;
  }

  async updateItem(id: number, name: string): Promise<Item> {
    const item = await this.getItemById(id);
    item.name = name;
    return this.itemsRepository.save(item);
  }

  async deleteItem(id: number): Promise<void> {
    const result = await this.itemsRepository.delete(id);
    if (result.affected === 0) {
      throw new NotFoundException('Item not found');
    }
  }
}
