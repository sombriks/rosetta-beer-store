package beer.store.model;

import java.util.Date;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import javax.persistence.Temporal;
import javax.persistence.TemporalType;
import javax.xml.bind.annotation.XmlAccessType;
import javax.xml.bind.annotation.XmlAccessorType;
import javax.xml.bind.annotation.XmlRootElement;

@Entity
@XmlRootElement
@Table(name = "beer")
@XmlAccessorType(XmlAccessType.FIELD)
public class Beer {

	@Id
	@Column(name = "idbeer")
	private long idBeer;

	@Temporal(TemporalType.TIMESTAMP)
	@Column(name = "creationdatebeer")
	private Date creationDateBeer;
}
