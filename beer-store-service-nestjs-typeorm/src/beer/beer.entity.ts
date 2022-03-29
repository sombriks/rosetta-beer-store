import {
  Entity,
  Column,
  PrimaryGeneratedColumn,
  CreateDateColumn,
} from 'typeorm';

@Entity()
export class Beer {
  @PrimaryGeneratedColumn()
  idbeer: number;
  @CreateDateColumn()
  creationdatebeer: Date;
  @Column()
  titlebeer: string;
  @Column()
  descriptionbeer: string;
  @Column()
  idmedia: number;
  // @OneToOne(() => Media)
  // @JoinColumn()
  // media: Media;
}
