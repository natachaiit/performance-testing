import {
  Controller,
  Get,
  Post,
  Body,
  Param,
  Put,
  Delete,
} from '@nestjs/common';
import { ItemsService } from './items.service';
import { Item } from './item.entity';

@Controller('items')
export class ItemsController {
  constructor(private readonly itemsService: ItemsService) {}

  @Post()
  createItem(@Body('name') name: string): Promise<Item> {
    return this.itemsService.createItem(name);
  }

  @Get()
  getItems(): Promise<Item[]> {
    return this.itemsService.getItems();
  }

  @Get(':id')
  getItem(@Param('id') id: number): Promise<Item> {
    return this.itemsService.getItemById(id);
  }

  @Put(':id')
  updateItem(
    @Param('id') id: number,
    @Body('name') name: string,
  ): Promise<Item> {
    return this.itemsService.updateItem(id, name);
  }

  @Delete(':id')
  deleteItem(@Param('id') id: number): Promise<void> {
    return this.itemsService.deleteItem(id);
  }
}
