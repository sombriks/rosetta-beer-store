package beer.store.model;

import java.util.Date;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;
import javax.persistence.Temporal;
import javax.persistence.TemporalType;

import lombok.Data;

@Data
@Entity
@Table(name = "beer")
public class Beer {

  @Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  @Column(name = "idbeer")
  private Long idBeer;

  @Temporal(TemporalType.TIMESTAMP)
  @Column(name = "creationdatebeer")
  private Date creationDateBeer;

  @Column(name = "titlebeer")
  private String titleBeer;

  @Column(name = "descriptionbeer")
  private String descriptionBeer;

  @Column(name="idmedia")
  private Long idMedia;

}