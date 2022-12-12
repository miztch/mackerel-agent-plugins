// Code generated by "tool/gen_mackerel_plugin.pl"; DO NOT EDIT
package main

import (
	"fmt"

	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-accesslog/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-apache2/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-aws-cloudfront/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-aws-dynamodb/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-aws-ec2-cpucredit/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-aws-ec2-ebs/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-aws-elasticache/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-aws-elasticsearch/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-aws-elb/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-aws-kinesis-streams/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-aws-lambda/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-aws-rds/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-aws-s3-requests/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-aws-ses/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-conntrack/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-docker/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-elasticsearch/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-fluentd/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-flume/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-gostats/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-graphite/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-h2o/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-haproxy/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-inode/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-jmx-jolokia/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-jvm/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-linux/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-mailq/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-memcached/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-mongodb/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-multicore/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-munin/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-mysql/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-nginx/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-openldap/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-php-apc/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-php-fpm/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-php-opcache/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-plack/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-postgres/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-proc-fd/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-rabbitmq/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-redis/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-sidekiq/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-snmp/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-solr/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-squid/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-td-table-count/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-trafficserver/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-twemproxy/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-unicorn/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-uptime/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-uwsgi-vassal/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-varnish/lib"
	"github.com/mackerelio/mackerel-agent-plugins/mackerel-plugin-xentop/lib"
)

func runPlugin(plug string) error {
	switch plug {
	case "accesslog":
		mpaccesslog.Do()
	case "apache2":
		mpapache2.Do()
	case "aws-cloudfront":
		mpawscloudfront.Do()
	case "aws-dynamodb":
		mpawsdynamodb.Do()
	case "aws-ec2-cpucredit":
		mpawsec2cpucredit.Do()
	case "aws-ec2-ebs":
		mpawsec2ebs.Do()
	case "aws-elasticache":
		mpawselasticache.Do()
	case "aws-elasticsearch":
		mpawselasticsearch.Do()
	case "aws-elb":
		mpawselb.Do()
	case "aws-kinesis-streams":
		mpawskinesisstreams.Do()
	case "aws-lambda":
		mpawslambda.Do()
	case "aws-rds":
		mpawsrds.Do()
	case "aws-s3-requests":
		mpawss3requests.Do()
	case "aws-ses":
		mpawsses.Do()
	case "conntrack":
		mpconntrack.Do()
	case "docker":
		mpdocker.Do()
	case "elasticsearch":
		mpelasticsearch.Do()
	case "fluentd":
		mpfluentd.Do()
	case "flume":
		mpflume.Do()
	case "gostats":
		mpgostats.Do()
	case "graphite":
		mpgraphite.Do()
	case "h2o":
		mph2o.Do()
	case "haproxy":
		mphaproxy.Do()
	case "inode":
		mpinode.Do()
	case "jmx-jolokia":
		mpjmxjolokia.Do()
	case "jvm":
		mpjvm.Do()
	case "linux":
		mplinux.Do()
	case "mailq":
		mpmailq.Do()
	case "memcached":
		mpmemcached.Do()
	case "mongodb":
		mpmongodb.Do()
	case "multicore":
		mpmulticore.Do()
	case "munin":
		mpmunin.Do()
	case "mysql":
		mpmysql.Do()
	case "nginx":
		mpnginx.Do()
	case "openldap":
		mpopenldap.Do()
	case "php-apc":
		mpphpapc.Do()
	case "php-fpm":
		mpphpfpm.Do()
	case "php-opcache":
		mpphpopcache.Do()
	case "plack":
		mpplack.Do()
	case "postgres":
		mppostgres.Do()
	case "proc-fd":
		mpprocfd.Do()
	case "rabbitmq":
		mprabbitmq.Do()
	case "redis":
		mpredis.Do()
	case "sidekiq":
		mpsidekiq.Do()
	case "snmp":
		mpsnmp.Do()
	case "solr":
		mpsolr.Do()
	case "squid":
		mpsquid.Do()
	case "td-table-count":
		mptdtablecount.Do()
	case "trafficserver":
		mptrafficserver.Do()
	case "twemproxy":
		mptwemproxy.Do()
	case "unicorn":
		mpunicorn.Do()
	case "uptime":
		mpuptime.Do()
	case "uwsgi-vassal":
		mpuwsgivassal.Do()
	case "varnish":
		mpvarnish.Do()
	case "xentop":
		mpxentop.Do()
	default:
		return fmt.Errorf("unknown plugin: %q", plug)
	}
	return nil
}

var plugins = []string{
	"accesslog",
	"apache2",
	"aws-cloudfront",
	"aws-dynamodb",
	"aws-ec2-cpucredit",
	"aws-ec2-ebs",
	"aws-elasticache",
	"aws-elasticsearch",
	"aws-elb",
	"aws-kinesis-streams",
	"aws-lambda",
	"aws-rds",
	"aws-s3-requests",
	"aws-ses",
	"conntrack",
	"docker",
	"elasticsearch",
	"fluentd",
	"flume",
	"gostats",
	"graphite",
	"h2o",
	"haproxy",
	"inode",
	"jmx-jolokia",
	"jvm",
	"linux",
	"mailq",
	"memcached",
	"mongodb",
	"multicore",
	"munin",
	"mysql",
	"nginx",
	"openldap",
	"php-apc",
	"php-fpm",
	"php-opcache",
	"plack",
	"postgres",
	"proc-fd",
	"rabbitmq",
	"redis",
	"sidekiq",
	"snmp",
	"solr",
	"squid",
	"td-table-count",
	"trafficserver",
	"twemproxy",
	"unicorn",
	"uptime",
	"uwsgi-vassal",
	"varnish",
	"xentop",
}
