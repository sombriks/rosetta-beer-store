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
import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;

@Entity
@XmlRootElement
@Table(name = "beer")
@XmlAccessorType(XmlAccessType.FIELD)
public class Beer {

	@Id
	@XmlElement(name = "idbeer")
	@Column(name = "idbeer")
	private long idBeer;

	@Temporal(TemporalType.TIMESTAMP)
	@Column(name = "creationdatebeer")
	private Date creationDateBeer;

	@Column(name = "titlebeer")
	private String titleBeer;

	@Column(name = "descriptionbeer")
	private String descriptionBeer;

	@Column(name = "idmedia")
	private Long idMedia;

	public long getIdBeer() {
		return idBeer;
	}

	public void setIdBeer(long idBeer) {
		this.idBeer = idBeer;
	}

	public Date getCreationDateBeer() {
		return creationDateBeer;
	}

	public void setCreationDateBeer(Date creationDateBeer) {
		this.creationDateBeer = creationDateBeer;
	}

	public String getTitleBeer() {
		return titleBeer;
	}

	public void setTitleBeer(String titleBeer) {
		this.titleBeer = titleBeer;
	}

	public String getDescriptionBeer() {
		return descriptionBeer;
	}

	public void setDescriptionBeer(String descriptionBeer) {
		this.descriptionBeer = descriptionBeer;
	}

	public Long getIdMedia() {
		return idMedia;
	}

	public void setIdMedia(Long idMedia) {
		this.idMedia = idMedia;
	}

}
