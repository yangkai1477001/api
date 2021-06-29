package v1alpha2


type SmfSvcCfg struct {
     // SmfPort default 80 \n
     // Optional
     SmfPort uint `mapstructure:"smfPort" json:"smfPort,omitempty"`

}

type SmfConfig struct { 
     // The major configurations of sm-svc
     // Mandatory.
     SvcCfg SmfSvcCfg `mapstructure:"svcCfg" json:"svcCfg"`
}
